package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadFromStdin() string {
	r := bufio.NewReader(os.Stdin)
	answer, _ := r.ReadString('\n')
	return strings.TrimSpace(answer)
}

func ReadCharFromStdin() string {
	r := bufio.NewReader(os.Stdin)
	rune, _, _ := r.ReadRune()
	return string(rune)
}
