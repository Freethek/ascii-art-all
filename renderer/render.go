package renderer

import (
	"ascii-art-all/color"
	"strings"
)

// The codes in this file goes to the render branch

func Render(ansiCode, substring, input string, bannerMap map[rune][]string) []string {
	splitedForNewline := strings.Split(input, "\\n")

	if len(splitedForNewline) > 0 && splitedForNewline[len(splitedForNewline)-1] == "" {
		splitedForNewline = splitedForNewline[:len(splitedForNewline)-1]
	}

	var result []string

	for _, segment := range splitedForNewline {
		if segment == "" {
			result = append(result, "")
		} else {
			//checking for subtring occurence in this segment
			toBeColored := color.FindColoredIndices(segment, substring)
			converted := []rune(segment)

			for row := 0; row <= 7; row++ {
				line := ""

				for i, ch := range converted {
					//entering the colored appending
					if ansiCode != "" && toBeColored[i] && (i == 0 || toBeColored[i-1] == false) {
						line += ansiCode
					}

					//the main graphic but may be colored due to the prefixed ascii code
					//it may also not be colored
					line += bannerMap[ch][row]

					//to know if we need to exit this block that appends asci code to suffix
					//so it dont spill over to the next word
					if ansiCode != "" && toBeColored[i] && (i == len(converted)-1 || !toBeColored[i+1]) {
						line += "\033[0m"
					}
				}

				result = append(result, line)
			}
		}
	}

	return result
}
