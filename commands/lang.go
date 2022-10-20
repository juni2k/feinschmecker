package commands

import (
	"github.com/juni2k/feinschmecker/config"
	"github.com/juni2k/feinschmecker/filter"
	"github.com/juni2k/feinschmecker/lang"
	"github.com/juni2k/feinschmecker/sessions"
)

// lang.go contains language-switching commands

func En(conf *config.Config, session *sessions.Session) *TextMap {
	session.Language = lang.En
	session.Save()

	return &TextMap{
		session.Language: filter.AddHeading(
			"Updated language preference.",
			"Für Deutsch bitte /de schicken.",
		),
	}
}

func De(conf *config.Config, session *sessions.Session) *TextMap {
	session.Language = lang.De
	session.Save()

	return &TextMap{
		session.Language: filter.AddHeading(
			"Sprache aktualisiert.",
			"Use /en to switch back to English.",
		),
	}
}
