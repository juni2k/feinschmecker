package reply

import (
	"log"

	"github.com/juni2k/feinschmecker/commands"
	"github.com/juni2k/feinschmecker/lang"
)

type Reply struct {
	TextMap *commands.TextMap
}

func (r *Reply) Translation(language lang.Language) string {
	text, ok := (*r.TextMap)[language]
	if !ok {
		log.Fatalf("Language %d not available in reply %+v", language, r)
	}

	return text
}
