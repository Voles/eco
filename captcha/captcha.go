// Package captcha wraps github.com/dchest/captcha and provides an sqlite store for it.
//
// Add to your router:
//
//	router.Handler("GET", "/captcha/:fn", captcha.Handler())
//
// Add to your html template:
//
//	.Parse(captcha.TemplateText)
//
// Pass captcha data to the template:
//
//	type MyTemplateData struct {
//	    Captcha captcha.TemplateData
//	    // ...
//
// Create a captcha in your GET handler:
//
//	myTemplateData.Captcha.ID = captcha.New()
//
// In your POST handler, call Verify after validating other input because Verify invalidates the captcha. If you're executing the template again, you must create a new captcha.
//
//	if !captcha.Verify(id, answer) {
//	    data.Captcha.ID = captcha.New()
//	    data.Captcha.Err = true
//	    html.MyTemplate.Execute(w, data)
//		return
//	}
package captcha

import (
	"net/http"

	"github.com/dchest/captcha"
	_ "github.com/mattn/go-sqlite3"
)

func Handler() http.Handler {
	return captcha.Server(captcha.StdWidth, captcha.StdHeight)
}

func New() (id string) {
	return captcha.NewLen(6)
}

// Verify checks and invalidates the captcha.
func Verify(id string, digits string) bool {
	return captcha.VerifyString(id, digits)
}
