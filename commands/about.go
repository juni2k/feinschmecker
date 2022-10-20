package commands

import (
	"github.com/juni2k/feinschmecker/config"
	"github.com/juni2k/feinschmecker/filter"
	"github.com/juni2k/feinschmecker/lang"
	"github.com/juni2k/feinschmecker/sessions"
)

// about.go contains the about command.

func About(conf *config.Config, session *sessions.Session) *TextMap {
	return &TextMap{
		lang.En: filter.AddHeading(
			"Key",
			filter.Perl("", "./icons.pl", "en"),
		),
		lang.De: filter.AddHeading(
			"Legende",
			filter.Perl("", "./icons.pl", "de"),
		),
	}
}
