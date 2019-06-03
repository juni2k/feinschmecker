package menu

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"

	"github.com/nanont/feinschmecker/filter"
)

type Request int

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
	Language string
	Year     uint
	Day      uint8
}

const (
	Now  Request = iota
	Next Request = iota
)

const (
	MenuURL     = "https://speiseplan.studierendenwerk-hamburg.de/en/570/2019/0/"
	MenuUrlTmpl = "https://speiseplan.studierendenwerk-hamburg.de/{{.Language}}/570/{{.Year}}/{{.Day}}/"
)

func Show(request Request) string {
	resp, err := http.Get(urlFor(request))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	menu := parse(urlFor(request), resp.Body)
	//	fmt.Printf("%+v", menu)

	template, err := template.ParseFiles("templates/menu.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = template.Execute(buf, menu)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func urlFor(request Request) string {
	params := UrlParams{"en", 2019, 0}

	if request == Now {
		params.Day = 0
	} else if request == Next {
		params.Day = 99
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

	//	fmt.Printf(string(site))
	doc, err := goquery.NewDocumentFromReader(siteBody)
	if err != nil {
		log.Fatal(err)
	}

	m.Date = doc.Find("tr#headline th.category").First().Text()
	m.Date = filter.Strip(m.Date)

	doc.Find("div#plan tr.odd, div#plan tr.even").Each(
		func(i int, dish *goquery.Selection) {
			d := Dish{}

			desc := dish.Find(".dish-description").First()

			d.Label = desc.Text()
			d.Label = filter.Perl(d.Label, "./label.pl")

			d.Price = dish.Find(".price").First().Text()
			d.Price = filter.Strip(d.Price)

			icons := desc.Find("img").Map(
				func(j int, img *goquery.Selection) string {
					attr, exists := img.Attr("alt")

					if exists {
						return attr
					} else {
						log.Fatalf("attr %s doesn't exist", attr)
						return ""
					}
				})
			d.Icons = strings.Join(icons, "\n")
			d.Icons = filter.Perl(d.Icons, "./icons.pl")

			m.Dishes = append(m.Dishes, d)
		})

	return m
}
