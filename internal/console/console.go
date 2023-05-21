package console

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Console struct{}

func NewConsole() (*Console, error) {
	return &Console{}, nil
}

// Input
// Read console input.
func (c *Console) Input() string {
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
