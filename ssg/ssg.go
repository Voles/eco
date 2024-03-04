// Package ssg creates both static and dynamic sites from translated HTML templates and markdown files.
//
// The content directory root may contain:
//
//   - html template files
//   - one folder for each html page, containing markdown files whose filename root is the language prefix, like "en.md"
//   - files and folders which are copied verbatim (see Keep)
//
// The output is like "/en/page.html".
//
// Note that "gotext update" requires a Go module and package for merging translations, accessing message.DefaultCatalog and writing catalog.go.
// While gotext-update-templates has been extended to accept additional directories, a root module and package is still required for static site generation.
//
// For symlink support see [Handler] and [StaticHTML].
package ssg

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	paths "path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/dys2p/eco/lang"
	"gitlab.com/golang-commonmark/markdown"
	"golang.org/x/text/language"
	_ "golang.org/x/text/message"
)

var Keep = []string{
	"ads.txt",
	"app-ads.txt",
	"assets",
	"files",
	"images",
	"sites",
	"static",
}

var md = markdown.New(markdown.HTML(true), markdown.Linkify(false))

// Language can be used in templates.
type Language struct {
	BCP47    string
	Name     string
	Prefix   string
	Selected bool
}

// SelectLanguage returns a [Language] slice. If if only one language is present, the slice will be empty.
func SelectLanguage(langs []lang.Lang, selected lang.Lang) []Language {
	var languages []Language
	if len(langs) > 1 {
		for _, l := range langs {
			languages = append(languages, Language{
				BCP47:    l.BCP47,
				Name:     strings.ToUpper(l.Prefix),
				Prefix:   l.Prefix,
				Selected: l.Prefix == selected.Prefix,
			})
		}
	}
	return languages
}

type TemplateData struct {
	lang.Lang
	Languages []Language // usually empty if only one language is defined
	Path      string     // without language prefix, use for language buttons and hreflang
}

// Hreflangs returns <link hreflang> elements for every td.Language, including the selected language.
//
// See also: https://developers.google.com/search/blog/2011/12/new-markup-for-multilingual-content
func (td TemplateData) Hreflangs() template.HTML {
	var b strings.Builder
	for _, l := range td.Languages {
		b.WriteString(fmt.Sprintf(`<link rel="alternate" hreflang="%s" href="/%s/%s">`, l.BCP47, l.Prefix, td.Path))
		b.WriteString("\n")
	}
	return template.HTML(b.String())
}

// Handler returns a HTTP handler which serves content from fsys.
// It accepts an additional HTML template and a function which makes custom template data.
// For compatibility with StaticHTML, the custom template data struct should embed TemplateData.
//
// Note that embed.FS does not support symlinks. If you use symlinks to share content,
// consider building a go:generate workflow which calls "cp --dereference".
func Handler(fsys fs.FS, add *template.Template, makeTemplateData func(*http.Request, TemplateData) any) (http.Handler, error) {
	handler := http.NewServeMux()
	err := generate(
		fsys,
		add,
		func(path string) {
			path = paths.Join("/", path)
			handler.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				http.ServeFileFS(w, r, fsys, path)
			})
		},
		func(path string, t *template.Template, data TemplateData) {
			path = paths.Join("/", path)
			handler.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				err := t.ExecuteTemplate(w, "html", makeTemplateData(r, data))
				if err != nil {
					log.Printf("error executing ssg template %s: %v", path, err)
				}
			})
		},
	)
	return handler, err
}

// ListenAndServe provides an easy way to preview a static site with absolute src and href paths.
func ListenAndServe(dir string) {
	log.Println("listening to 127.0.0.1:8080")
	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.ListenAndServe("127.0.0.1:8080", nil)
}

// StaticHTML creates static HTML files. Templates are executed with TemplateData. Symlinks are dereferenced.
//
// # Example
//
//	package main
//
//	import "github.com/dys2p/eco/ssg"
//
//	//go:generate gotext-update-templates -srclang=en-US -lang=de-DE,en-US -out=catalog.go -d . -d ./example.com
//
//	func main() {
//		ssg.StaticHTML("./example.com", "/tmp/build/example.com")
//		ssg.ListenAndServe("/tmp/build/example.com")
//	}
func StaticHTML(srcDir, outDir string) {
	if realOutDir, err := filepath.EvalSymlinks(outDir); err == nil {
		outDir = realOutDir
	}
	if !strings.HasPrefix(outDir, "/tmp/") {
		log.Fatalf("refusing to write outside of /tmp")
	}
	_ = os.RemoveAll(outDir)

	err := generate(
		os.DirFS(srcDir),
		nil,
		func(path string) {
			src := filepath.Join(srcDir, path)
			dstDir := filepath.Dir(filepath.Join(outDir, path))
			if err := os.MkdirAll(dstDir, 0700); err != nil {
				log.Fatalf("error making folder %s: %v", dstDir, err)
			}
			err := exec.Command("cp", "--archive", "--dereference", "--target-directory", dstDir, src).Run()
			if err != nil {
				log.Fatalf("error copying %s to %s: %v", src, dstDir, err)
			}
		},
		func(path string, t *template.Template, data TemplateData) {
			dst := filepath.Join(outDir, path)
			if err := os.MkdirAll(filepath.Dir(dst), 0700); err != nil {
				log.Fatalf("error making folder %s: %v", filepath.Dir(dst), err)
			}
			outfile, err := os.Create(dst)
			if err != nil {
				log.Fatalf("error opening outfile %s: %v", dst, err)
			}
			defer outfile.Close()
			err = t.ExecuteTemplate(outfile, "html", data)
			if err != nil {
				log.Fatalf("error executing template for %s: %v", dst, err)
			}
		},
	)
	if err != nil {
		log.Fatalf("error %v", err)
	}
}

func generate(fsys fs.FS, add *template.Template, serve func(path string), execute func(path string, t *template.Template, data TemplateData)) error {
	langs := lang.MakeLanguages(nil, "de", "en")

	// serve static content and collect sites
	var sites []string
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return fmt.Errorf("reading root dir: %w", err)
	}
	for _, entry := range entries {
		var isDir = entry.IsDir()
		// support symlink to dir
		if entry.Type()&fs.ModeSymlink != 0 {
			info, _ := fs.Stat(fsys, entry.Name()) // fs.Stat returns symlink target FileInfo
			if info.Mode()&fs.ModeDir != 0 {
				isDir = true
			}
		}

		switch {
		case strings.HasPrefix(entry.Name(), "."):
			continue
		case slices.Contains(Keep, entry.Name()):
			serve(entry.Name())
		case isDir:
			sites = append(sites, entry.Name())
		}
	}

	// prepare site template
	tmpl, err := template.ParseFS(fsys, "*.html")
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}
	if add != nil {
		for _, t := range add.Templates() {
			if t.Tree != nil { // that's possible
				tmpl, err = tmpl.AddParseTree(t.Name(), t.Tree)
				if err != nil {
					return fmt.Errorf("adding additional template %s: %w", t.Name(), err)
				}
			}
		}
	}

	// translate sites
	for _, site := range sites {

		// read markdown files
		var bcp47 []string
		var content []string // same indices
		entries, err := fs.ReadDir(fsys, site)
		if err != nil {
			return fmt.Errorf("reading dir %s: %w", site, err)
		}
		for _, entry := range entries {
			if strings.HasPrefix(entry.Name(), ".") {
				continue
			}
			if entry.IsDir() {
				continue
			}
			ext := filepath.Ext(entry.Name())
			root := strings.TrimSuffix(entry.Name(), ext)
			if ext == ".html" || ext == ".md" {
				filecontent, err := fs.ReadFile(fsys, filepath.Join(site, entry.Name()))
				if err != nil {
					return fmt.Errorf("reading file: %w", err)
				}
				if ext == ".md" {
					filecontent = []byte(md.RenderToString(filecontent))
				}
				bcp47 = append(bcp47, root)
				content = append(content, string(filecontent))
			}
		}
		if len(content) == 0 {
			continue
		}

		// make matcher for available translations
		var haveTags []language.Tag
		for _, have := range bcp47 {
			haveTags = append(haveTags, language.MustParse(have))
		}
		matcher := language.NewMatcher(haveTags)

		// assemble site template
		for _, lang := range langs {
			_, index, _ := matcher.Match(lang.Tag)
			tt, err := tmpl.Clone()
			if err != nil {
				return fmt.Errorf("cloning template: %w", err)
			}
			tt, err = tt.Parse(`{{define "content"}}` + content[index] + `{{end}}`) // or parse content into t and then call AddParseTree(content, t.Tree)
			if err != nil {
				return fmt.Errorf("executing site %s: %w", site, err)
			}
			outpath := filepath.Join(lang.Prefix, site+".html")
			execute(outpath, tt, TemplateData{
				Lang:      lang,
				Languages: SelectLanguage(langs, lang),
				Path:      site + ".html",
			})
		}
	}
	return nil
}