package commands

import (
	"github.com/nanont/feinschmecker/filter"
	"github.com/nanont/feinschmecker/lang"
	"github.com/nanont/feinschmecker/reply"
	"github.com/nanont/feinschmecker/sessions"
)

// about.go contains the about command.

func About(session *sessions.Session) *reply.Reply {
	return &reply.Reply{TextMap: map[lang.Language]string{
		lang.En: filter.AddHeading("Key",
			filter.Perl("", "./icons.pl", "en")),
		lang.De: filter.AddHeading("Legende",
			filter.Perl("", "./icons.pl", "de")),
	}}
}
