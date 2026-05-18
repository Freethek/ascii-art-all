package output

import (
	"ascii-art-all/color"
	"os"
	"strings"
)

func WriteToFile(filename string, lines []string) error {
	result := ""
	var stripped []string

	for _, line := range lines {
		stripped = append(stripped, color.StripAnsi(line))
	}

	result = strings.Join(stripped, "\n")

	err := os.WriteFile(filename, []byte(result), 0644)

	return err
}
