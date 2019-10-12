package commands

import (
	"github.com/nanont/feinschmecker/lang"
	"github.com/nanont/feinschmecker/reply"
	"github.com/nanont/feinschmecker/sessions"
)

// lang.go contains language-switching commands

func En(session *sessions.Session) *reply.Reply {
	session.Language = lang.En
	session.Save()

	return &reply.Reply{TextMap: map[lang.Language]string{
		session.Language: "Excellent!",
	}}
}

func De(session *sessions.Session) *reply.Reply {
	session.Language = lang.De
	session.Save()

	return &reply.Reply{TextMap: map[lang.Language]string{
		session.Language: "Sehr gut!",
	}}
}
