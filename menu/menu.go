package menu

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"text/template"

	"github.com/juni2k/feinschmecker/bindata"
	"github.com/juni2k/feinschmecker/filter"
	"github.com/juni2k/feinschmecker/lang"

	"github.com/PuerkitoBio/goquery"
)

type Request int
type Language int

type Menu struct {
	Date   string
	Link   string
	Dishes []Dish
}

type Dish struct {
	Label string
	Price string
	Icons string
}

type UrlParams struct {
	Location     int
	LanguagePath string
	Day          string
}

const (
	Now  Request = iota
	Next Request = iota
)

const (
	CanteenID   = 158
	MenuUrlTmpl = "https://www.stwhh.de/{{.LanguagePath}}?l={{.Location}}&t={{.Day}}"
)

func Show(request Request, language lang.Language) string {
	resp, err := http.Get(urlFor(request, language))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	menu := parse(urlFor(request, language), resp.Body)

	tmplBytes, err := bindata.Asset("templates/menu.txt")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("menu").Parse(string(tmplBytes))
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, menu)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func urlFor(request Request, language lang.Language) string {
	params := UrlParams{Location: CanteenID}

	if request == Now {
		params.Day = "today"
	} else if request == Next {
		params.Day = "next_day"
	}

	if language == lang.En {
		params.LanguagePath = "en/menu"
	} else if language == lang.De {
		params.LanguagePath = "speiseplan"
	}

	tmpl, err := template.New("menuUrl").Parse(MenuUrlTmpl)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, params)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func parse(url string, siteBody io.ReadCloser) *Menu {
	m := &Menu{}

	m.Link = url

	doc, err := goquery.NewDocumentFromReader(siteBody)
	if err != nil {
		log.Fatal(err)
	}

	// Fuck this shit entirely. It used to be much easier. This will probably break.
	m.Date = doc.Find(".button__text--highlight").First().Text()
	m.Date = filter.Strip(m.Date)

	// STWHH returns data for all the canteens, but it is hidden, except for the one we requested.
	// Are they out of their minds?
	iHopeSomeoneGetsFiredOverThis := fmt.Sprintf(
		"div.tx-epwerkmenu-menu-location-container[data-location-id=\"%d\"]",
		CanteenID)
	theCanteenWeReallyCareAboutWithoutOtherGarbage := doc.Find(iHopeSomeoneGetsFiredOverThis)
	theCanteenWeReallyCareAboutWithoutOtherGarbage.Find("div.singlemeal").Each(
		func(i int, dish *goquery.Selection) {
			d := Dish{}

			desc := dish.Find("h5.singlemeal__headline").First()

			d.Label = desc.Text()
			d.Label = filter.Perl(d.Label, "./label.pl")

			// This used to be .price, but now we need to reach deep into the HTML and do bullshit regex heuristics.
			somewhereDownHere := dish.Find(".singlemeal__bottom").Text()
			matches := regexp.MustCompile("(\\d,\\d{2} €)\\s*(Studierende|Students)").FindStringSubmatch(somewhereDownHere)
			if len(matches) > 2 {
				d.Price = filter.Strip(matches[1])
			} else {
				d.Price = "[could not extract price]"
			}

			// There used to be semantically valid alt-texts. Not anymore. Hope you're not blind.
			icons := dish.Find(".singlemeal__icontooltip").Map(
				func(j int, tooltip *goquery.Selection) string {
					attr, exists := tooltip.Attr("title")

					if exists {
						// Parse the icon title out of the tooltip
						matches := regexp.MustCompile("<b>(.*?)</b>").FindStringSubmatch(attr)
						if len(matches) > 1 {
							return matches[1]
						} else {
							html, _ := tooltip.Html()
							log.Printf("could not extract title from tooltip: %s", html)
							return ""
						}
					} else {
						log.Fatalf("attr %s doesn't exist", attr)
						return ""
					}
				},
			)
			d.Icons = strings.Join(icons, "\n")
			d.Icons = filter.Perl(d.Icons, "./icons.pl")

			m.Dishes = append(m.Dishes, d)
		},
	)

	return m
}
