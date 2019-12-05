package commands

import (
	"github.com/nanont/feinschmecker/config"
	"github.com/nanont/feinschmecker/reply"
	"github.com/nanont/feinschmecker/sessions"
)

type Func func(*config.Config, *sessions.Session) *reply.Reply
type Map map[string]Func
