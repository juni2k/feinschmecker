package filter

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Strip(s string) string {
	return strings.TrimSpace(s)
}

func Perl(in string, scriptName string) string {
	selfDir := filepath.Dir(os.Args[0])

	cmd := exec.Command(scriptName)
	cmd.Dir = filepath.Join(selfDir, "perl")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	_, err = stdin.Write([]byte(in))
	if err != nil {
		log.Fatal(err)
	}
	stdin.Close()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	// cmd.Dir =
	return string(out)
}
