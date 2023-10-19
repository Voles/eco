// Package countries contains European countries and their VAT rates from an European Union point of view.
package countries

import (
	"sort"

	"github.com/dys2p/eco/lang"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type Country string // ISO 3166-1 code

const (
	AT Country = "AT"
	BE Country = "BE"
	BG Country = "BG"
	CY Country = "CY"
	CZ Country = "CZ"
	DE Country = "DE"
	DK Country = "DK"
	EE Country = "EE"
	ES Country = "ES"
	FI Country = "FI"
	FR Country = "FR"
	GR Country = "GR"
	HR Country = "HR"
	HU Country = "HU"
	IE Country = "IE"
	IT Country = "IT"
	LT Country = "LT"
	LU Country = "LU"
	LV Country = "LV"
	MT Country = "MT"
	NL Country = "NL"
	PL Country = "PL"
	PT Country = "PT"
	RO Country = "RO"
	SE Country = "SE"
	SI Country = "SI"
	SK Country = "SK"

	NonEU Country = "non-EU"

	CH Country = "CH"
	GB Country = "GB"
)

var EuropeanUnion = []Country{AT, BE, BG, CY, CZ, DE, DK, EE, ES, FI, FR, GR, HR, HU, IE, IT, LT, LU, LV, MT, NL, PL, PT, RO, SE, SI, SK}

func (c Country) TranslateName(langstr string) string {
	l := lang.Lang(langstr)
	switch c {
	case AT:
		return l.Tr("Austria")
	case BE:
		return l.Tr("Belgium")
	case BG:
		return l.Tr("Bulgaria")
	case CH:
		return l.Tr("Switzerland")
	case CY:
		return l.Tr("Cyprus")
	case CZ:
		return l.Tr("Czechia")
	case DE:
		return l.Tr("Germany")
	case DK:
		return l.Tr("Denmark")
	case EE:
		return l.Tr("Estonia")
	case ES:
		return l.Tr("Spain")
	case FI:
		return l.Tr("Finland")
	case FR:
		return l.Tr("France")
	case GB:
		return l.Tr("United Kingdom")
	case GR:
		return l.Tr("Greece")
	case HR:
		return l.Tr("Croatia")
	case HU:
		return l.Tr("Hungary")
	case IE:
		return l.Tr("Ireland")
	case IT:
		return l.Tr("Italy")
	case LT:
		return l.Tr("Lithuania")
	case LU:
		return l.Tr("Luxembourg")
	case LV:
		return l.Tr("Latvia")
	case MT:
		return l.Tr("Malta")
	case NL:
		return l.Tr("Netherlands")
	case PL:
		return l.Tr("Poland")
	case PT:
		return l.Tr("Portugal")
	case RO:
		return l.Tr("Romania")
	case SE:
		return l.Tr("Sweden")
	case SI:
		return l.Tr("Slovenia")
	case SK:
		return l.Tr("Slovakia")
	default:
		return ""
	}
}

// VAT rates are from: https://europa.eu/youreurope/business/taxation/vat/vat-rules-rates/index_en.htm#shortcut-5.
// Note that the ISO 3166-1 code for Greece is "GR", but the VAT rate table uses its ISO 639-1 code "EL".
func (c Country) VAT() VATRates {
	switch c {
	case AT:
		return map[string]float64{
			"standard":  0.20,
			"reduced-1": 0.10,
			"reduced-2": 0.13,
			"parking":   0.13,
		}
	case BE:
		return map[string]float64{
			"standard":  0.21,
			"reduced-1": 0.06,
			"reduced-2": 0.12,
			"parking":   0.12,
		}
	case BG:
		return map[string]float64{
			"standard":  0.20,
			"reduced-1": 0.09,
		}
	case CY:
		return map[string]float64{
			"standard":  0.19,
			"reduced-1": 0.05,
			"reduced-2": 0.09,
		}
	case CZ:
		return map[string]float64{
			"standard":  0.21,
			"reduced-1": 0.10,
			"reduced-2": 0.15,
		}
	case DE:
		return map[string]float64{
			"standard":  0.19,
			"reduced-1": 0.07,
		}
	case DK:
		return map[string]float64{
			"standard": 0.25,
		}
	case EE:
		return map[string]float64{
			"standard":  0.20,
			"reduced-1": 0.09,
		}
	case ES:
		return map[string]float64{
			"standard":      0.21,
			"reduced-1":     0.10,
			"super-reduced": 0.04,
		}
	case FI:
		return map[string]float64{
			"standard":  0.24,
			"reduced-1": 0.10,
			"reduced-2": 0.14,
		}
	case FR:
		return map[string]float64{
			"standard":      0.20,
			"reduced-1":     0.055,
			"reduced-2":     0.10,
			"super-reduced": 0.021,
		}
	case GR:
		return map[string]float64{
			"standard":  0.24,
			"reduced-1": 0.06,
			"reduced-2": 0.13,
		}
	case HR:
		return map[string]float64{
			"standard":  0.25,
			"reduced-1": 0.05,
			"reduced-2": 0.13,
		}
	case HU:
		return map[string]float64{
			"standard":  0.27,
			"reduced-1": 0.05,
			"reduced-2": 0.18,
		}
	case IE:
		return map[string]float64{
			"standard":      0.23,
			"reduced-1":     0.09,
			"reduced-2":     0.135,
			"super-reduced": 0.048,
			"parking":       0.135,
		}
	case IT:
		return map[string]float64{
			"standard":      0.22,
			"reduced-1":     0.05,
			"reduced-2":     0.10,
			"super-reduced": 0.04,
		}
	case LT:
		return map[string]float64{
			"standard":  0.21,
			"reduced-1": 0.05,
			"reduced-2": 0.09,
		}
	case LU:
		return map[string]float64{
			"standard":      0.17,
			"reduced-1":     0.08,
			"super-reduced": 0.03,
			"parking":       0.14,
		}
	case LV:
		return map[string]float64{
			"standard":  0.21,
			"reduced-1": 0.12,
			"reduced-2": 0.05,
		}
	case MT:
		return map[string]float64{
			"standard":  0.18,
			"reduced-1": 0.05,
			"reduced-2": 0.07,
		}
	case NL:
		return map[string]float64{
			"standard":  0.21,
			"reduced-1": 0.09,
		}
	case PL:
		return map[string]float64{
			"standard":  0.23,
			"reduced-1": 0.05,
			"reduced-2": 0.08,
		}
	case PT:
		return map[string]float64{
			"standard":  0.23,
			"reduced-1": 0.06,
			"reduced-2": 0.13,
			"parking":   0.13,
		}
	case RO:
		return map[string]float64{
			"standard":  0.19,
			"reduced-1": 0.05,
			"reduced-2": 0.09,
		}
	case SE:
		return map[string]float64{
			"standard":  0.25,
			"reduced-1": 0.06,
			"reduced-2": 0.12,
		}
	case SI:
		return map[string]float64{
			"standard":  0.22,
			"reduced-1": 0.05,
			"reduced-2": 0.095,
		}
	case SK:
		return map[string]float64{
			"standard":  0.20,
			"reduced-1": 0.10,
		}
	default:
		return nil
	}
}

type CountryWithName struct {
	Country
	Name string
}

func TranslateAndSort(langstr string, countries []Country) []CountryWithName {
	var result = make([]CountryWithName, len(countries))
	for i := range countries {
		result[i] = CountryWithName{
			Country: countries[i],
			Name:    countries[i].TranslateName(langstr),
		}
	}

	collator := collate.New(language.Make(langstr), collate.IgnoreCase)
	sort.Slice(result, func(i, j int) bool {
		return collator.CompareString(result[i].Name, result[j].Name) < 0
	})
	return result
}
