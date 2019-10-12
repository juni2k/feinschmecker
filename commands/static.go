package commands

import (
	"bytes"
	"github.com/nanont/feinschmecker/bindata"
	"github.com/nanont/feinschmecker/config"
	"github.com/nanont/feinschmecker/lang"
	"github.com/nanont/feinschmecker/reply"
	"github.com/nanont/feinschmecker/sessions"
	"log"
	"text/template"
)

// static.go contains static commands
// that don't fetch resources etc. and
// always return the same text.

func Default(conf *config.Config, session *sessions.Session) *reply.Reply {
	return &reply.Reply{TextMap: map[lang.Language]string{
		lang.En: "I was unable to understand you :/",
		lang.De: "Ich konnte dich leider nicht verstehen :/",
	}}
}

func Start(conf *config.Config, session *sessions.Session) *reply.Reply {
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

	return &reply.Reply{TextMap: map[lang.Language]string{
		session.Language: buf.String(),
	}}
}
