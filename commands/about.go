package commands

import (
	"github.com/nanont/feinschmecker/config"
	"github.com/nanont/feinschmecker/filter"
	"github.com/nanont/feinschmecker/lang"
	"github.com/nanont/feinschmecker/sessions"
)

// about.go contains the about command.

func About(conf *config.Config, session *sessions.Session) *TextMap {
	return &TextMap{
		lang.En: filter.AddHeading("Key",
			filter.Perl("", "./icons.pl", "en")),
		lang.De: filter.AddHeading("Legende",
			filter.Perl("", "./icons.pl", "de")),
	}
}
