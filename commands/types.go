package commands

import (
	"github.com/nanont/feinschmecker/config"
	"github.com/nanont/feinschmecker/lang"
	"github.com/nanont/feinschmecker/sessions"
)

type Func func(*config.Config, *sessions.Session) *TextMap
type Map map[string]Func
type TextMap map[lang.Language]string
