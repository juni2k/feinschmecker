package commands

import (
	"github.com/nanont/feinschmecker/lang"
	"github.com/nanont/feinschmecker/reply"
	"github.com/nanont/feinschmecker/sessions"
)

// static.go contains static commands
// that don't fetch resources etc. and
// always return the same text.

func Default(session *sessions.Session) *reply.Reply {
	return &reply.Reply{TextMap: map[lang.Language]string{
		lang.En: "I was unable to understand you :/",
		lang.De: "Ich konnte dich leider nicht verstehen :/",
	}}
}

func Start(session *sessions.Session) *reply.Reply {
	return &reply.Reply{TextMap: map[lang.Language]string{
		lang.En: "This is the /start stub.",
		lang.De: "Dies ist ein Platzhalter für /start.",
	}}
}
