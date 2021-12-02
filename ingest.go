package common

import (
	"bufio"
	"bytes"
	"os"
)

// Ingest reads a file in to memory.
func Ingest(name string) ([]byte, error) {
	return os.ReadFile(name)
}

// SplitLines returns a line-by-line string slice of the input []byte(s).
func SplitLines(b ...[]byte) ([][]string, error) {
	lines := make([][]string, len(b))
	for i := range b {
		s := bufio.NewScanner(bytes.NewReader(b[i]))
		for s.Scan() {
			lines[i] = append(lines[i], s.Text())
		}
	}
	return lines, nil
}
