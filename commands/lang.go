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

func Apply(conf *config.Config, session *sessions.Session) *TextMap {
	details := `Hi there!

As I am no longer a student at TUHH, I would like to see some person adopt this bot (maintenance, hosting, etc.)

The good:
- The code base is readable, clean and relatively mature.
- It's written in Go, so you can already prepare for a soul-crushing job at some of Hamburg's hippest dotcom start-ups! Taste the Freedom!
- You take over a real project with real responsibility instead of jacking off to fetish porn in your step-parents' basement!

The bad:
- You'd have to be responsible for the bot.

The ugly:
- That's you!

Interested? Write to @bananont.
The bot is going to be taken offline on July 15 if you don't.`

	return &TextMap {
		session.Language: filter.AddHeading(
			"Mensa bot takeover details",
			details),
	}
}
