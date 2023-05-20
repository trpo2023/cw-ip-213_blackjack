package cmd

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Cmd struct{}

func NewCmd() *Cmd {
	return &Cmd{}
}

// Input
// Read console input.
func (c *Cmd) Input() string {
	reader := bufio.NewReader(os.Stdin)

	if str, _, err := reader.ReadLine(); err != nil {
		if err != io.EOF {
			log.Printf("\nerror reading user input: %v", err)
		}
		return ""
	} else {
		return strings.TrimRight(string(str), "\r\n")
	}
}
