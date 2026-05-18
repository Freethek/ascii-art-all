package main

import (
	"ascii-art-all/banner"
	"ascii-art-all/color"
	"ascii-art-all/output"
	"ascii-art-all/renderer"
	"fmt"
	"os"
	"strings"
)

func main() {
	//creating align and banner variable
	alignType := "left"
	bannerName := "standard"
	var input, substring, colorName, outputFileName string

	//to store the none flag value
	var positional []string

	//looping through the argument to identify flags and get their values
	for _, args := range os.Args[1:] {
		if strings.HasPrefix(args, "--") {
			//we want to check if there is an align flag first
			if strings.HasPrefix(args, "--align=") {
				//getting the align value
				alignType = validateAlignType(strings.TrimPrefix(args, "--align="))
			} else if strings.HasPrefix(args, "--color=") {
				//getting the color value
				colorName = strings.TrimPrefix(args, "--color=")
			} else if strings.HasPrefix(args, "--output=") {
				//getting the output value
				outputFileName = strings.TrimPrefix(args, "--output=")
			} else {
				fmt.Fprintf(os.Stderr, "wrong banner flag")
				Error()
			}
		} else {
			//getting the none flag output,
			//which are the input, substring, and the bannername
			//maiximum of three values are expected here
			//minimum of 1 value is expected
			positional = append(positional, args)
		}
	}

	//the flags present in the argument will determine the way i will assign my input, substring and banner from the positionals
	//i need to know what the user wants from the arguments parsed
	//if there is no colorname i should not be expecting a substring but an just an input and a banner
	//so if colorname is not empty, positionals can not be empty too nether will the values exceed three
	//if there is a color flag and there is no substring we will just stil color the whole input
	if colorName != "" && len(positional) != 0 && len(positional) <= 3 {
		if len(positional) == 1 {
			input = positional[0]
		} else if len(positional) == 2 {
			//since there is a color flag
			//its possible that the second positional value can either be a substring or a banner
			//so i need to check if the second is a validbannerfirst else i will take it as a substring
			if isValidBanner(positional[1]) {
				input = positional[0]
				bannerName = positional[1]
			} else {
				substring = positional[0]
				input = positional[1]
			}
		} else {
			substring = positional[0]
			input = positional[1]
			if isValidBanner(positional[2]) {
				bannerName = positional[2]
			} else {
				fmt.Fprintf(os.Stderr, "The Banner Provided Is Invalid")
				Error()
			}
		}

	} else {
		// if there is no color flag i dont want a substring
		//i should work with only two positionals
		if len(positional) == 1 {
			input = positional[0]
		} else if len(positional) == 2 {
			if isValidBanner(positional[1]) {
				input = positional[0]
				bannerName = positional[1]
			} else {
				Error()
			}
		} else {
			Error()
		}
	}
	//handling empty input
	if input == "" {
		fmt.Println()
		os.Exit(0)
	}

	//getting the ansicode of the color name
	ansiCode := ""
	if colorName != "" {
		ansiCode = color.GetAnsiCode(colorName)
		if ansiCode == "" {
			fmt.Fprintf(os.Stderr, "Invalid color: %v\n", colorName)
			Error()
		}
	}

	//offloading the ascii art from a file and loading it into the banner map
	bannerMap := banner.Load("banners/" + bannerName + ".txt")

	//getting the ascii art to be rendered baseed on the input
	toBerendered := renderer.Render(ansiCode, substring, input, bannerMap)

	//validate the output filename is having the right extension before writing the art to the file
	if outputFileName != "" {
		if !strings.HasSuffix(outputFileName, ".txt") {
			fmt.Fprintf(os.Stderr, "The Output file name must end with .txt extension\n")
			os.Exit(1)
		}
		err := output.WriteToFile(outputFileName, toBerendered)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Writing Into File %v\n", err)
			os.Exit(1)
		}
	}

	//getting the current terminal width
	getTerminal := GetWidth()

	//printing all that n
	result := Align(toBerendered, alignType, input, getTerminal, bannerMap, ansiCode, substring)

	for _, line := range result {
		fmt.Println(line)
	}

}

// printing the error message
func Error() {
	fmt.Fprintf(os.Stderr, "Usage: go run . [OPTIONS(align,color,output)]... [SUBSTRING(not compulsory)] [STRING] [BANNER(not compulsory)]\n\nExample: go run . --align=right something standard")
	os.Exit(1)
}

// validating the align type
func validateAlignType(alignType string) string {
	validTypes := map[string]bool{
		"left": true, "right": true, "center": true, "justify": true,
	}

	if !validTypes[alignType] {
		fmt.Fprintf(os.Stderr, "Wrong align type value")
		os.Exit(1)
	}

	return alignType
}

// validating banner type
func isValidBanner(bannerName string) bool {
	_, err := os.Stat("banners/" + bannerName + ".txt")
	return err == nil
}

//stoped  at swam land
