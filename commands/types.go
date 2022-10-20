package commands

import (
	"github.com/juni2k/feinschmecker/config"
	"github.com/juni2k/feinschmecker/lang"
	"github.com/juni2k/feinschmecker/sessions"
)

type Func func(*config.Config, *sessions.Session) *TextMap
type Map map[string]Func
type TextMap map[lang.Language]string
