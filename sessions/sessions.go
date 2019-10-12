package sessions

import (
	"encoding/json"
	"fmt"
	"github.com/nanont/feinschmecker/lang"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Session struct {
	ID       int64
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

	sessions := make(SessionMap)

	/* Populate session map from cache, if possible */
	sessionFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, sessionFile := range sessionFiles {
		absPath := filepath.Join(dir, sessionFile.Name())
		log.Printf("Loading session %s", absPath)

		jsonBytes, err := ioutil.ReadFile(absPath)
		if err != nil {
			log.Fatal(err)
		}

		var session Session

		err = json.Unmarshal(jsonBytes, &session)
		if err != nil {
			log.Fatal(err)
		}

		sessions[session.ID] = &session
	}

	return sessions
}

func New(id int64) *Session {
	return &Session{ID: id, Language: lang.En}
}

func GetOrNew(sessions SessionMap, id int64) *Session {
	session, ok := sessions[id]
	if !ok {
		session = New(id)
		sessions[id] = session
	}

	return session
}

func (s *Session) Save() {
	outPath := filepath.Join(dir, fmt.Sprintf("%d.json", s.ID))

	jsonBytes, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(outPath, jsonBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
