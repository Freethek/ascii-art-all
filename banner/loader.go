package banner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Load(filename string) map[rune][]string {
	// open the file
	fileOpened, err := os.Open(filename)
	// checking if there is an error when opening the filename
	if err != nil {
		fmt.Fprintf(os.Stderr, "there was an error reading the banner file %v\n", err)
		os.Exit(1)
	}
	// close the opened file after reading it to save memory space
	defer fileOpened.Close()

	// create a slice to save lines which stores every line from the file, recall we got 855lines from the txt file.
	var lines []string
	// scanner helps read the file content step by step
	scanner := bufio.NewScanner(fileOpened)
	// reading the file line by line
	for scanner.Scan() {
		line := scanner.Text()               // it get the current line as a text
		line = strings.TrimRight(line, "\r") // to remove the /r
		lines = append(lines, line)          // adding the line to the lines slice
	}

	// create a map to find the AScii art of each character
	bannerMap := make(map[rune][]string) // note "make" helps create a new map ready for use, without the "make" the map is nil
	// and will crash  if we try to  add data.

	// loop through the ASCII  charaters
	for i := 32; i <= 126; i++ { // it loops from ASCII 32 - 126
		start := (i-32)*9 + 1 // it finds where each character start from
		//getting the 8-lines of the ASCII ART
		grab8lines := lines[start : start+8] // it takes 8-lines from the slice which represent one  charater
		// storing the character in the map
		bannerMap[rune(i)] = grab8lines
	}
	return bannerMap
}
