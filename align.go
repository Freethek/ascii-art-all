package main

import (
	"ascii-art-all/color"
	"ascii-art-all/renderer"
	"strings"
)

func Align(lines []string, alignType, input string, termwidth int, bannerMap map[rune][]string, ansiCode, substring string) []string {
	var result []string

	//we want to split the space evenly between words in the terminal
	if alignType == "justify" {
		cleanInput := strings.ReplaceAll(input, "\\n", " ")
		words := strings.Fields(cleanInput)

		//if the input is one word then we just render it like its center
		if len(words) == 1 {
			for _, line := range lines {
				padding := (termwidth - color.VisibleLen(line)) / 2
				result = append(result, strings.Repeat(" ", padding)+line)
			}
		} else {
			var renderedWords [][]string
			//collecting all the rendered words first
			for _, word := range words {
				renderedWord := renderer.Render(ansiCode, substring, word, bannerMap)
				renderedWords = append(renderedWords, renderedWord)
			}
			//calculating the aggregate width of each word
			totalWidth := 0
			for _, row := range renderedWords {
				totalWidth += color.VisibleLen(row[0])
			}
			//calculating spaces to be added in between words during rendering
			gaps := len(words) - 1
			spacePerGap := (termwidth - totalWidth) / gaps
			for i := 0; i <= 7; i++ {
				row := ""
				for j, rw := range renderedWords {
					row += rw[i]
					//helps to avoid appending the spaces to the last word
					if j < len(renderedWords)-1 {
						row += strings.Repeat(" ", spacePerGap)
					}
				}
				result = append(result, row)
			}
		}
		return result
	}

	for _, line := range lines {
		switch alignType {
		case "left":
			//default mode
			result = append(result, line)
		case "right":
			//getting the amount of spaces to be preappended
			padding := termwidth - color.VisibleLen(line)
			result = append(result, strings.Repeat(" ", padding)+line)
		case "center":
			padding := (termwidth - color.VisibleLen(line)) / 2
			result = append(result, strings.Repeat(" ", padding)+line)
		}
	}
	return result
}
