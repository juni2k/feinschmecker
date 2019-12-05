package commands

import (
	"github.com/nanont/feinschmecker/config"
	"github.com/nanont/feinschmecker/filter"
	"github.com/nanont/feinschmecker/lang"
	"github.com/nanont/feinschmecker/sessions"
)

// lang.go contains language-switching commands

func En(conf *config.Config, session *sessions.Session) *TextMap {
	session.Language = lang.En
	session.Save()

	return &TextMap{
		session.Language: filter.AddHeading(
			"Updated language preference.",
			"Für Deutsch bitte /de schicken."),
	}
}

func De(conf *config.Config, session *sessions.Session) *TextMap {
	session.Language = lang.De
	session.Save()

	return &TextMap{
		session.Language: filter.AddHeading(
			"Sprache aktualisiert.",
			"Use /en to switch back to English."),
	}
}
