package commands

import (
	"bytes"
	"log"
	"text/template"

	"github.com/juni2k/feinschmecker/bindata"
	"github.com/juni2k/feinschmecker/config"
	"github.com/juni2k/feinschmecker/lang"
	"github.com/juni2k/feinschmecker/sessions"
)

// static.go contains static commands
// that don't fetch resources etc. and
// always return the same text.

func Default(conf *config.Config, session *sessions.Session) *TextMap {
	return &TextMap{
		lang.En: "I was unable to understand you :/",
		lang.De: "Ich konnte dich leider nicht verstehen :/",
	}
}

func Start(conf *config.Config, session *sessions.Session) *TextMap {
	var tmplPath string
	if session.Language == lang.En {
		tmplPath = "templates/start.en.txt"
	} else if session.Language == lang.De {
		tmplPath = "templates/start.de.txt"
	}

	tmplBytes, err := bindata.Asset(tmplPath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("start").Parse(string(tmplBytes))
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, conf)
	if err != nil {
		log.Fatal(err)
	}

	return &TextMap{
		session.Language: buf.String(),
	}
}
