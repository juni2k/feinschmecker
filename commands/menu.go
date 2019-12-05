package commands

import (
	"github.com/nanont/feinschmecker/config"
	"github.com/nanont/feinschmecker/menu"
	"github.com/nanont/feinschmecker/sessions"
)

// menu.go contains commands for
// fetching available menus

func Now(conf *config.Config, session *sessions.Session) *TextMap {
	return &TextMap{
		session.Language: menu.Show(menu.Now, session.Language),
	}
}

func Next(conf *config.Config, session *sessions.Session) *TextMap {
	return &TextMap{
		session.Language: menu.Show(menu.Next, session.Language),
	}
}
