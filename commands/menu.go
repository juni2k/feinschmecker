package commands

import (
	"github.com/nanont/feinschmecker/config"
	"github.com/nanont/feinschmecker/lang"
	"github.com/nanont/feinschmecker/menu"
	"github.com/nanont/feinschmecker/reply"
	"github.com/nanont/feinschmecker/sessions"
)

// menu.go contains commands for
// fetching available menus

func Now(conf *config.Config, session *sessions.Session) *reply.Reply {
	return &reply.Reply{TextMap: map[lang.Language]string{
		session.Language: menu.Show(menu.Now, session.Language),
	}}
}

func Next(conf *config.Config, session *sessions.Session) *reply.Reply {
	return &reply.Reply{TextMap: map[lang.Language]string{
		session.Language: menu.Show(menu.Next, session.Language),
	}}
}
