package utils

import (
	"fmt"
	"log"
	"os/exec"
)

func CrashWithError(message string) {
	exec.Command("say", fmt.Sprintf("'%s'", message))
	log.Fatal(message)
}
