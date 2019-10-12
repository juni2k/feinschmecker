package filter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Strip(s string) string {
	return strings.TrimSpace(s)
}

func AddHeading(heading string, text string) string {
	return strings.Join([]string{
		fmt.Sprintf("*%s*", heading),
		text,
	}, "\n\n")
}

func Perl(in string, scriptName string, args ...string) string {
	selfDir := filepath.Dir(os.Args[0])

	cmd := exec.Command(scriptName, args...)
	cmd.Dir = filepath.Join(selfDir, "perl")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	_, err = stdin.Write([]byte(in))
	if err != nil {
		log.Fatal(err)
	}
	err = stdin.Close()
	if err != nil {
		log.Fatal(err)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
