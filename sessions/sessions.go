package sessions

import (
	"github.com/nanont/feinschmecker/lang"
	"log"
	"os"
	"path/filepath"
)

type Session struct {
	Language lang.Language
}

type SessionMap map[int64]*Session

// Directory where sessions are kept in
var dir string

func Init(workdir string) SessionMap {
	dir = filepath.Join(workdir, "sessions")
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	/* TODO: init from cache etc. */
	sessions := make(SessionMap)
	return sessions
}

func New() *Session {
	return &Session{}
}

func GetOrNew(sessions SessionMap, id int64) *Session {
	session, ok := sessions[id]
	if !ok {
		session = New()
		sessions[id] = session
	}

	return session
}
