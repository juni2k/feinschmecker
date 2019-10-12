package reply

import (
	"github.com/nanont/feinschmecker/lang"
	"log"
)

type Reply struct {
	TextMap map[lang.Language]string
}

func (r *Reply) Translation(language lang.Language) string {
	text, ok := r.TextMap[language]
	if !ok {
		log.Fatalf("Language %d not available in reply %+v", language, r)
	}

	return text
}


