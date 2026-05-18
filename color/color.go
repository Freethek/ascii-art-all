package color

import "strings"

var ansiColors = map[string]string{
	// Standard colors
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	// Bright colors
	"bright-black":   "\033[90m",
	"bright-red":     "\033[91m",
	"bright-green":   "\033[92m",
	"bright-yellow":  "\033[93m",
	"bright-blue":    "\033[94m",
	"bright-magenta": "\033[95m",
	"bright-cyan":    "\033[96m",
	"bright-white":   "\033[97m",
	// Background colors
	"bg-black":   "\033[40m",
	"bg-red":     "\033[41m",
	"bg-green":   "\033[42m",
	"bg-yellow":  "\033[43m",
	"bg-blue":    "\033[44m",
	"bg-magenta": "\033[45m",
	"bg-cyan":    "\033[46m",
	"bg-white":   "\033[47m",
	// Bright background colors
	"bg-bright-black":   "\033[100m",
	"bg-bright-red":     "\033[101m",
	"bg-bright-green":   "\033[102m",
	"bg-bright-yellow":  "\033[103m",
	"bg-bright-blue":    "\033[104m",
	"bg-bright-magenta": "\033[105m",
	"bg-bright-cyan":    "\033[106m",
	"bg-bright-white":   "\033[107m",
	// Text styles
	"bold":          "\033[1m",
	"dim":           "\033[2m",
	"italic":        "\033[3m",
	"underline":     "\033[4m",
	"blink":         "\033[5m",
	"reverse":       "\033[7m",
	"strikethrough": "\033[9m",

	// Extended RGB colors - Foreground
	"orange":     "\033[38;2;255;165;0m",
	"pink":       "\033[38;2;255;192;203m",
	"purple":     "\033[38;2;128;0;128m",
	"brown":      "\033[38;2;139;69;19m",
	"gold":       "\033[38;2;255;215;0m",
	"silver":     "\033[38;2;192;192;192m",
	"coral":      "\033[38;2;255;127;80m",
	"salmon":     "\033[38;2;250;128;114m",
	"lime":       "\033[38;2;0;255;0m",
	"teal":       "\033[38;2;0;128;128m",
	"navy":       "\033[38;2;0;0;128m",
	"maroon":     "\033[38;2;128;0;0m",
	"olive":      "\033[38;2;128;128;0m",
	"indigo":     "\033[38;2;75;0;130m",
	"violet":     "\033[38;2;238;130;238m",
	"lavender":   "\033[38;2;230;230;250m",
	"turquoise":  "\033[38;2;64;224;208m",
	"tan":        "\033[38;2;210;180;140m",
	"khaki":      "\033[38;2;240;230;140m",
	"crimson":    "\033[38;2;220;20;60m",
	"beige":      "\033[38;2;245;245;220m",
	"mint":       "\033[38;2;152;255;152m",
	"ivory":      "\033[38;2;255;255;240m",
	"charcoal":   "\033[38;2;54;69;79m",
	"rose":       "\033[38;2;255;0;127m",
	"peach":      "\033[38;2;255;218;185m",
	"amber":      "\033[38;2;255;191;0m",
	"emerald":    "\033[38;2;0;201;87m",
	"ruby":       "\033[38;2;155;17;30m",
	"sapphire":   "\033[38;2;15;82;186m",
	"topaz":      "\033[38;2;255;200;124m",
	"jade":       "\033[38;2;0;168;107m",
	"scarlet":    "\033[38;2;255;36;0m",
	"bronze":     "\033[38;2;205;127;50m",
	"copper":     "\033[38;2;184;115;51m",
	"plum":       "\033[38;2;142;69;133m",
	"mauve":      "\033[38;2;224;176;255m",
	"periwinkle": "\033[38;2;204;204;255m",
	"aqua":       "\033[38;2;0;255;255m",
	"cream":      "\033[38;2;255;253;208m",

	// Extended RGB colors - Background
	"bg-orange":    "\033[48;2;255;165;0m",
	"bg-pink":      "\033[48;2;255;192;203m",
	"bg-purple":    "\033[48;2;128;0;128m",
	"bg-brown":     "\033[48;2;139;69;19m",
	"bg-gold":      "\033[48;2;255;215;0m",
	"bg-silver":    "\033[48;2;192;192;192m",
	"bg-coral":     "\033[48;2;255;127;80m",
	"bg-salmon":    "\033[48;2;250;128;114m",
	"bg-lime":      "\033[48;2;0;255;0m",
	"bg-teal":      "\033[48;2;0;128;128m",
	"bg-navy":      "\033[48;2;0;0;128m",
	"bg-maroon":    "\033[48;2;128;0;0m",
	"bg-olive":     "\033[48;2;128;128;0m",
	"bg-indigo":    "\033[48;2;75;0;130m",
	"bg-violet":    "\033[48;2;238;130;238m",
	"bg-lavender":  "\033[48;2;230;230;250m",
	"bg-turquoise": "\033[48;2;64;224;208m",
	"bg-tan":       "\033[48;2;210;180;140m",
	"bg-khaki":     "\033[48;2;240;230;140m",
	"bg-crimson":   "\033[48;2;220;20;60m",
	"bg-beige":     "\033[48;2;245;245;220m",
	"bg-mint":      "\033[48;2;152;255;152m",
	"bg-rose":      "\033[48;2;255;0;127m",
	"bg-peach":     "\033[48;2;255;218;185m",
	"bg-amber":     "\033[48;2;255;191;0m",
	"bg-emerald":   "\033[48;2;0;201;87m",
	"bg-ruby":      "\033[48;2;155;17;30m",
	"bg-sapphire":  "\033[48;2;15;82;186m",
	"bg-scarlet":   "\033[48;2;255;36;0m",
	"bg-bronze":    "\033[48;2;205;127;50m",
	"bg-copper":    "\033[48;2;184;115;51m",
	"bg-plum":      "\033[48;2;142;69;133m",
	"bg-aqua":      "\033[48;2;0;255;255m",
	"bg-jade":      "\033[48;2;0;168;107m",
}

// getting the color code of color from the ansicode map
func GetAnsiCode(colorName string) string {
	AnsiCode, exists := ansiColors[colorName]

	if exists {
		return AnsiCode
	} else {
		return ""
	}
}

// checking if the substrings is present in the input
func FindColoredIndices(input, substring string) map[int]bool {
	//declaring the variable it be returned which holds the index of the character to colored in the input
	indexes := make(map[int]bool)

	if input == "" {
		return indexes
	}

	//marking all the indices as true
	if substring == "" {
		for i := range []rune(input) {
			indexes[i] = true
		}
	} else {
		start := 0
		for {
			//checking for the start position of the found substring in the input
			ind := strings.Index(input[start:], substring)
			if ind == -1 {
				break
			}
			//we need to mark all the substring true using their index position in the input
			//but the index will be stored in the map int of  bool
			for j := 0; j < len([]rune(substring)); j++ {
				indexes[start+ind+j] = true
			}
			//this enable the next iteration to run using the previous index as the new start index
			//to enable the function search for another occurence in the substring
			start += ind + len([]rune(substring))
		}

	}

	return indexes
}



//stripping the ansi code from the ascii art that will be rendered
//this is important because we dont want the apended ansi codes at the begining and end of the ascci art to be counted 
//this will give acuraccy to the padding which is the added space 
func StripAnsi(s string) string {
	//rebuilding the string by appending only the ascii art
	//exempting the color codes
	//to get the accurate count of the exact of the ascii itself
	result := ""
	inEscape := false

	//looping through the ascii art and looking for escape sequence
	for _, ch := range s {
		//this helps me identify ansi code at the end of the string so  can avoid adding it to the result
		if inEscape {
			if ch == 'm' {
				inEscape = false
			}
		} else {
			//this shows that there is a color code 
			// appended to the ascii art that needs to be striped out from the string
			if ch == '\033' {
				//at this point we are in the ansi code
				inEscape = true
			} else {
				//at this point we get only the art itself
				result += string(ch)
			}
		}
	}

	return result
}

//this will return the acurate count of the real ascii from the string returned
//from the stripedansi function 
//to help us get a real count when rendering justified colored ascii art 
func VisibleLen(s string) int {
	return len(StripAnsi(s))
}
