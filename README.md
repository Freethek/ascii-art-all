# ASCII Art All

A comprehensive command-line tool written in **Go** that transforms text into striking ASCII art with advanced rendering capabilities. This project combines four progressive modules — base ASCII rendering, color support, file output, and smart terminal-aware alignment — into a single, feature-rich application.

**Key Features:**
- 🎨 Render text as large ASCII art using 90+ banner fonts
- 🌈 Color individual substrings or entire output with 40+ color options
- 📄 Save output directly to files
- ↔️ Smart alignment (left, right, center, justify) with dynamic terminal width detection
- ✅ Comprehensive test coverage with unit tests
- 📦 Zero external dependencies — uses only Go standard library

---

## Table of Contents

- [Overview](#overview)
- [Core Architecture](#core-architecture)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
  - [Basic Rendering](#basic-rendering)
  - [Color Support](#color-support)
  - [File Output](#file-output)
  - [Alignment & Terminal Sizing](#alignment--terminal-sizing)
  - [Combined Options](#combined-options)
- [Options Reference](#options-reference)
  - [--align](#--align)
  - [--color](#--color)
  - [--output](#--output)
- [Banner Fonts](#banner-fonts)
- [Color Palettes](#color-palettes)
- [Examples](#examples)
- [Architecture & Implementation](#architecture--implementation)
- [Unit Testing](#unit-testing)
- [Edge Cases & Handling](#edge-cases--handling)
- [Design Principles](#design-principles)

---

## Overview

**ASCII Art All** is a unified implementation of multiple ASCII art projects, combining the following requirements into one powerful tool:

1. **ascii-art (Base)**: Renders text as ASCII art using banner files
2. **ascii-art-color (Layer 1)**: Adds color support for substrings or entire output
3. **ascii-art-output (Layer 2)**: Enables saving rendered output to files
4. **ascii-art-justify (Layer 3)**: Provides multiple alignment modes with terminal-aware sizing

The application features modular architecture with dedicated packages for each concern (banner loading, color management, rendering, output, and alignment), making it maintainable, testable, and extensible.

---

## Core Architecture

### Design Philosophy

The project follows **single responsibility principle** with clear separation of concerns:

- **`banner/`** — Handles banner file loading and character-to-ASCII-art mapping
- **`color/`** — Manages ANSI color codes, color validation, and substring highlighting
- **`renderer/`** — Orchestrates the rendering of input text into ASCII art with optional coloring
- **`output/`** — Handles file I/O operations for saving rendered output
- **`align.go`** — Implements alignment logic (left, right, center, justify)
- **`terminal.go`** — Detects terminal width via syscalls for responsive design
- **`main.go`** — Command-line argument parsing and orchestrates the pipeline

### Pipeline Flow

```
Command-line Args
       ↓
Argument Parser (main.go)
       ↓
Extract: --align, --color, --output, [options]
       ↓
Load Banner File (banner/loader.go)
       ↓
Render ASCII Art (renderer/render.go)
       ├─→ Apply color to substrings (color/color.go)
       └─→ Generate 8-line ASCII output
       ↓
Apply Alignment (align.go)
├─→ Detect terminal width (terminal.go)
└─→ Apply left/right/center/justify
       ↓
Output (dual path)
├─→ Print to stdout
└─→ Write to file if --output flag (output/output.go)
```

---

## Project Structure

```
ascii-art-all/
├── main.go                           # CLI entry point, flag parsing, pipeline orchestration
├── go.mod                            # Go module definition
├── align.go                          # Alignment logic: left, right, center, justify
├── terminal.go                       # Terminal width detection (Linux/Unix syscalls)
│
├── banner/                           # Banner file loading and parsing
│   ├── loader.go                     # Parse .txt banner files into map[rune][]string
│   └── loader_test.go                # Unit tests for banner loading
│
├── color/                            # Color management and ANSI codes
│   ├── color.go                      # Color palette, substring matching, ANSI utilities
│   └── (40+ color options supported)
│
├── renderer/                         # ASCII art rendering engine
│   └── render.go                     # Render text to ASCII art with optional coloring
│
├── output/                           # File output operations
│   └── output.go                     # Write rendered ASCII art to .txt files
│
├── banners/                          # 90+ ASCII art banner fonts
│   ├── standard.txt
│   ├── shadow.txt
│   ├── thinkertoy.txt
│   ├── (87 additional banner variants)
│   └── ...
│
└── README.md                         # This file
```

---

## Installation

### Prerequisites

- **Go 1.22 or higher** installed on your system
- Linux/Unix environment (terminal width detection uses Linux syscalls)
- No external dependencies required

### Setup Steps

```bash
# Clone the repository
git clone https://acad.learn2earn.ng/git/your-username/ascii-art-all.git
cd ascii-art-all

# Verify the module
go mod tidy

# Build the executable (optional)
go build -o ascii-art

# Or run directly
go run . "hello"
```

---

## Usage

### Command Format

```bash
go run . [OPTIONS]... [SUBSTRING(optional)] [STRING] [BANNER(optional)]
```

### Basic Rendering

**Render text with default settings (left alignment, standard banner):**

```bash
go run . "hello"
```

**Output:**
```
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
```

**Render with a specific banner:**

```bash
go run . "hello" shadow
```

### Color Support

**Color entire output:**

```bash
go run . --color=red "hello"
```

**Color a specific substring:**

```bash
go run . --color=blue "kit" "a king kitten have kit"
```

The substrings "kit" in both locations will be colored blue while the rest remains uncolored.

**Color with a specific banner:**

```bash
go run . --color=emerald "test" "hello world" standard
```

### File Output

**Save rendered output to a .txt file:**

```bash
go run . --output=banner.txt "hello" standard
```

**The file must have a `.txt` extension:**

```bash
go run . --output=result.txt "hello"  # ✓ Valid
go run . --output=result "hello"      # ✗ Error: missing .txt
```

### Alignment & Terminal Sizing

**Center-align output:**

```bash
go run . --align=center "hello"
```

**Right-align output:**

```bash
go run . --align=right "hello"
```

**Justify output (evenly distributed across terminal width):**

```bash
go run . --align=justify "hello world" standard
```

**The alignment automatically adapts to your terminal width.** Resize your terminal and run the command again to see the output reflow.

### Combined Options

Most powerful — combine all features:

```bash
go run . --align=center --color=gold --output=art.txt "hello world" standard
```

This command will:
1. Render "hello world" using the standard banner
2. Color the entire output gold
3. Center-align it based on terminal width
4. Save the result (without ANSI codes) to `art.txt`
5. Also display it on stdout

---

## Options Reference

### --align

Controls text alignment within the terminal width.

| Value | Behavior |
|-------|----------|
| `left` | Text starts at the left edge (default) |
| `right` | Text ends at the right edge |
| `center` | Text is centered horizontally |
| `justify` | Text is evenly distributed left-to-right; words are spaced to fit full width |

**Example:**
```bash
go run . --align=center "hello"
go run . --align=justify "hello world" standard
```

**Terminal-aware:** All alignment modes adapt to your current terminal width. The alignment is recalculated based on `GetWidth()` which uses the `TIOCGWINSZ` ioctl.

### --color

Applies ANSI color codes to the output.

**Format:**
```bash
go run . --color=<colorname> [substring] <string> [banner]
```

- If `substring` is provided, only that substring is colored
- If no `substring` provided, the entire string is colored
- The substring can appear multiple times; all occurrences are colored

**Supported Colors:**
See [Color Palettes](#color-palettes) section below.

**Example:**
```bash
go run . --color=red "hello"              # Entire "hello" is red
go run . --color=blue "ll" "hello"        # Only "ll" in "hello" is blue
go run . --color=green "a" "banana" shadow  # All "a"s in "banana" are green
```

### --output

Saves the rendered ASCII art to a file.

**Format:**
```bash
go run . --output=<filename.txt> <string> [banner]
```

**Constraint:** The filename must end with `.txt` extension.

**Behavior:**
- ANSI color codes are stripped from the output file (plain text only)
- The file is created/overwritten with 0644 permissions
- Output is also displayed on stdout

**Example:**
```bash
go run . --output=result.txt "hello"
go run . --output=banner.txt "hello world" shadow
```

---

## Banner Fonts

The project includes **90+ ASCII art banner fonts**, each providing a unique style. Each banner file contains ASCII representations for characters 32–126 (printable ASCII range).

### Popular Banners

| Banner Name | Style | Usage |
|---|---|---|
| `standard` | Clean, professional | `go run . "hello" standard` |
| `shadow` | Shadowed effect | `go run . "hello" shadow` |
| `thinkertoy` | Blocky, geometric | `go run . "hello" thinkertoy` |
| `big` | Bold, large | `go run . "hello" big` |
| `block` | Solid blocks | `go run . "hello" block` |
| `bubble` | Rounded bubbles | `go run . "hello" bubble` |
| `digital` | Digital display style | `go run . "hello" digital` |

### Complete List

All 90+ banners are located in the `banners/` directory. Try any `.txt` filename (without extension) as your banner parameter:

```bash
ls banners/ | wc -l  # Lists all available banners
```

**Some notable variants:**
- 3D styles: `3d`, `3D-ASCII`, `3D Diagonal`, `3d_diagonal`
- AMC styles: `amcrazor`, `amcthin`, `amcneko`, `amcslash`
- Broadway styles: `Broadway`, `broadway_kb`
- And many more...

---

## Color Palettes

The `color/` package provides **40+ built-in colors** using ANSI escape codes:

### Standard Colors (8)
`black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`

### Bright Colors (8)
`bright-black`, `bright-red`, `bright-green`, `bright-yellow`, `bright-blue`, `bright-magenta`, `bright-cyan`, `bright-white`

### Background Colors
Standard: `bg-black`, `bg-red`, `bg-green`, ... (8 variants)  
Bright: `bg-bright-black`, `bg-bright-red`, ... (8 variants)

### Extended RGB Colors (24+)
`orange`, `pink`, `purple`, `brown`, `gold`, `silver`, `coral`, `salmon`, `lime`, `teal`, `navy`, `maroon`, `olive`, `indigo`, `violet`, `lavender`, `turquoise`, `tan`, `khaki`, `crimson`, `beige`, `mint`, `ivory`, `charcoal`, `rose`, `peach`, `amber`, `emerald`, `ruby`, `sapphire`, `topaz`, `jade`, `scarlet`, `bronze`, `copper`, `plum`, `mauve`, `periwinkle`, `aqua`, `cream`

### Text Styles
`bold`, `dim`, `italic`, `underline`, `blink`, `reverse`, `strikethrough`

### Color Code System

The project uses **ANSI 256-color and Truecolor (24-bit RGB)** escape sequences:

```
Standard:     \033[30m (e.g., black)
Bright:       \033[90m (e.g., bright-black)
Background:   \033[40m (e.g., bg-black)
RGB Extended: \033[38;2;255;165;0m (e.g., orange)
```

**Example Usage:**
```bash
go run . --color=scarlet "danger"
go run . --color=emerald "safe" standard
go run . --color=bright-cyan "hello"
go run . --color=bg-navy "text" > output.txt
```

---

## Examples

### Example 1: Basic Rendering
```bash
$ go run . "Hello"
```

### Example 2: Color with Substring Highlighting
```bash
$ go run . --color=magenta "or" "Hello World"

# Only "or" in "World" gets colored magenta
```

### Example 3: Center-Aligned with Custom Banner
```bash
$ go run . --align=center "ASCII" shadow
```

### Example 4: Save to File with Full Styling
```bash
$ go run . --align=right --color=gold --output=banner.txt "Success" standard
$ cat banner.txt
```

### Example 5: Justify Multiple Words
```bash
$ go run . --align=justify "hello world" standard
```
Words are spaced to fill the terminal width.

### Example 6: Newline Handling
```bash
$ go run . "Hello\nWorld"
```
Renders "Hello" and "World" on separate ASCII art blocks (with blank lines between).

### Example 7: Empty Input
```bash
$ go run . ""
# Outputs a single blank line
```

### Example 8: Multi-Word Color Application
```bash
$ go run . --color=red "kit" "a king kitten have kit"

# Both occurrences of "kit" (in "kitten" and standalone) are colored red
```

---

## Architecture & Implementation

### Module Breakdown

#### `banner/loader.go`
**Purpose:** Parse ASCII art banner files and create a lookup map.

**Key Logic:**
1. Open banner `.txt` file
2. Read all lines using bufio scanner
3. Extract 8-line blocks for each ASCII character (32–126)
4. Store in map with rune keys: `map[rune][]string`
5. Return map for fast O(1) character lookups during rendering

**Character Layout:**
```
ASCII 32 (' '):   lines 1–8     (offset = (32-32)*9+1 = 1)
ASCII 33 ('!'):   lines 10–17   (offset = (33-32)*9+1 = 10)
ASCII 34 ('"'):   lines 19–26   (offset = (34-32)*9+1 = 19)
...
ASCII 126 ('~'):  lines 855–862 (offset = (126-32)*9+1 = 855)
```

#### `color/color.go`
**Purpose:** Manage color codes and substring matching for selective coloring.

**Key Functions:**
- `GetAnsiCode(colorName string) string` — Lookup ANSI escape code by color name
- `FindColoredIndices(input, substring string) map[int]bool` — Find all positions of substring in input
- `StripAnsi(s string) string` — Remove ANSI codes from string (for file output)
- `VisibleLen(s string) int` — Get visible character count (excluding ANSI codes)

**Color Storage:**
All 40+ colors are stored in `ansiColors` map with their corresponding ANSI escape sequences.

#### `renderer/render.go`
**Purpose:** Orchestrate rendering of text to ASCII art with optional coloring.

**Algorithm:**
1. Split input by `\n` into segments
2. For each segment:
   - Create map of character indices to be colored (using `FindColoredIndices`)
   - Loop through 8 rows (height of each character)
   - For each character in segment:
     - Insert ANSI color code if character starts coloring
     - Append ASCII art from banner map
     - Insert reset code `\033[0m` if character ends coloring
3. Return array of 8 lines per segment

**Key Line:**
```go
line += bannerMap[ch][row]  // Append ASCII art for character at given row
```

#### `align.go`
**Purpose:** Apply alignment transformations based on terminal width.

**Alignment Modes:**
1. **Left:** No padding (default)
2. **Right:** Pad left side to push text right
3. **Center:** Pad left side by half of remaining width
4. **Justify:** Distribute words evenly across full width (complex math for multi-word input)

**Key Function:**
```go
padding := termwidth - VisibleLen(line)  // Account for invisible ANSI codes
```

#### `terminal.go`
**Purpose:** Detect terminal width dynamically.

**Mechanism:**
Uses Linux/Unix syscall `TIOCGWINSZ` (Terminal I/O Control: Get Window Size) to query terminal dimensions.

```go
syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), uintptr(syscall.TIOCGWINSZ), ...)
```

Extracts `ws.col` (column width) from the returned winsize struct.

#### `output/output.go`
**Purpose:** Write rendered ASCII art to files.

**Process:**
1. Strip ANSI codes from all lines (for clean text file)
2. Join lines with `\n`
3. Write to file with 0644 permissions

**Example Output File:**
```
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
```

#### `main.go`
**Purpose:** CLI argument parsing and pipeline orchestration.

**Argument Parsing Logic:**
1. Loop through `os.Args[1:]`
2. Extract flags: `--align=`, `--color=`, `--output=`
3. Collect positional arguments
4. Smart inference of `substring`, `input`, `bannerName` based on counts and validation
5. Validate banner exists, color valid, align type valid, output ends with `.txt`

**Pipeline:**
```
Parse args → Load banner → Render ASCII art → Apply alignment → Output (stdout + file)
```

---

## Unit Testing

The project includes unit tests for critical components.

### Running Tests

```bash
# Run all tests in the project
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test ./banner
```

### Test Files

#### `banner/loader_test.go`

Tests the banner loading functionality:

```go
func TestLoad(t *testing.T) {
    bannerMap := Load("../banners/standard.txt")
    
    // Verify all 95 ASCII characters loaded
    if len(bannerMap) != 95 {
        t.Errorf("incomplete graphic representation for all ascii")
    }
    
    // Verify each character has 8 lines
    if len(bannerMap['A']) != 8 {
        t.Errorf("invalid graphic representation")
    }
}
```

**What It Verifies:**
- ✓ Banner file loads successfully
- ✓ Exactly 95 characters parsed (ASCII 32–126)
- ✓ Each character has exactly 8 lines of ASCII art

### Test Coverage

Recommended areas for additional tests:
- Alignment calculations for different terminal widths
- Color substring matching (edge cases, overlaps)
- Newline handling in input
- Empty string and special character rendering
- File I/O error handling

---

## Edge Cases & Handling

The implementation handles numerous edge cases robustly:

### 1. Empty Input
```bash
$ go run . ""
# Output: Single blank line
$ go run . "\n"
# Output: Two blank lines
```

### 2. Newline Handling
```bash
$ go run . "Hello\nWorld"
# Renders Hello and World as separate ASCII blocks with blank lines between
```

**Implementation:** Split input by `\n` literal, render each segment independently.

### 3. Special Characters
```bash
$ go run . "!@#$%"
# All printable ASCII characters (32–126) are supported
```

### 4. Color Non-Matching Substring
```bash
$ go run . --color=red "xyz" "hello"
# No color applied; no "xyz" substring in "hello"
```

### 5. Multiple Color Occurrences
```bash
$ go run . --color=blue "l" "hello"
# Both 'l' characters colored blue
```

### 6. Alignment on Small Terminal
If terminal width is smaller than ASCII art width, output may overflow. The implementation doesn't truncate; it relies on terminal handling.

### 7. Invalid Banner
```bash
$ go run . "hello" nonexistent
# Error: "The Banner Provided Is Invalid"
```

### 8. Invalid Color
```bash
$ go run . --color=notacolor "hello"
# Error: "Invalid color: notacolor"
```

### 9. Invalid Alignment
```bash
$ go run . --align=diagonal "hello"
# Error: "Wrong align type value"
```

### 10. Invalid File Output Extension
```bash
$ go run . --output=banner "hello"
# Error: "The Output file name must end with .txt extension"
```

### 11. ANSI Code Length in Alignment Calculations
The alignment logic uses `VisibleLen()` to calculate width while accounting for invisible ANSI color codes, preventing misalignment.

---

## Design Principles

### 1. **Modularity**
Each package has a single, well-defined responsibility:
- `banner/` handles I/O and parsing
- `color/` manages color system
- `renderer/` orchestrates rendering
- `output/` handles file operations

### 2. **Composition over Inheritance**
Functions are composed to build complex behavior:
```
Render(colored ASCII) → Align(based on width) → Write(to stdout/file)
```

### 3. **Error Handling**
All errors exit with status code 1 and descriptive stderr messages:
```go
if err != nil {
    fmt.Fprintf(os.Stderr, "Error message: %v\n", err)
    os.Exit(1)
}
```

### 4. **Terminal-Aware Design**
Dynamic terminal width detection ensures output adapts to current window size.

### 5. **No External Dependencies**
Uses only Go standard library: `os`, `fmt`, `strings`, `bufio`, `syscall`.

### 6. **Testability**
Pure functions (like `Load()`, `FindColoredIndices()`) are easily tested in isolation.

### 7. **Go Best Practices**
- Proper error handling (checking `err != nil`)
- Idiomatic naming (CamelCase, short variable names)
- Deferred resource cleanup (`defer fileOpened.Close()`)
- Efficient data structures (maps for O(1) lookups)

---

## Command-Line Interface

### Usage Message

When invalid arguments are provided:

```
Usage: go run . [OPTIONS(align,color,output)]... [SUBSTRING(not compulsory)] [STRING] [BANNER(not compulsory)]

Example: go run . --align=right something standard
```

### Argument Parsing Strategy

The CLI intelligently handles flexible argument order:

```bash
# All valid:
go run . "hello"
go run . "hello" standard
go run . --align=center "hello"
go run . --color=red "hello"
go run . --color=red substring "hello" standard
go run . --align=center --color=red --output=file.txt "hello" shadow
```

**Parsing Logic:**
1. Prefixed with `--` → Flag (extract key-value)
2. Otherwise → Positional argument
3. Disambiguate positional args based on count and validation

---

## Performance Characteristics

- **Banner Loading:** O(1) per character lookup after initial O(n) file parse
- **Rendering:** O(m × h) where m = message length, h = height (8 lines)
- **Alignment:** O(n) where n = number of output lines
- **Color Matching:** O(m × s) where m = message length, s = substring length (worst case)

**Overall:** Linear in input size and terminal width. Suitable for interactive CLI use.

---

## Future Enhancements

Potential improvements for future versions:

1. **Windows Support** — Extend terminal width detection to Windows console API
2. **Custom Fonts** — Allow users to provide custom banner files
3. **Animation** — Add support for sequential rendering with delays
4. **Unicode Support** — Extend beyond ASCII 32–126 to support Unicode characters
5. **Theme System** — Predefined color schemes (dark mode, neon, pastel, etc.)
6. **Configuration File** — Allow default settings via `.ascii-artrc`
7. **Height Customization** — Support variable character heights beyond fixed 8 lines
8. **Performance Caching** — Cache frequently used banner-color combinations

---

## Troubleshooting

| Issue | Solution |
|-------|----------|
| Output overflows terminal | Terminal width detection may fail; try maximizing window |
| Color not applied | Verify color name exists; check spelling |
| File not created | Ensure filename ends with `.txt`; check write permissions |
| Banner not found | Use valid banner name; list with `ls banners/` |
| Alignment looks wrong | Terminal width may have changed; resize and re-run |

---

## References

- **GO Modules:** https://golang.org/ref/spec
- **ANSI Escape Codes:** https://en.wikipedia.org/wiki/ANSI_escape_code
- **Terminal I/O Control:** Linux man page `ioctl(2)`, `tty_ioctl(4)`
- **ASCII Ranges:** ASCII 32–126 are printable characters

---

## License & Attribution

This project demonstrates Go proficiency including file I/O, data structures, syscalls, and CLI design.

Developed as part of educational coursework combining four progressive ASCII art modules.
|---|---|
| `left` | Output is aligned to the left edge of the terminal (default) |
| `right` | Output is aligned to the right edge of the terminal |
| `center` | Output is centered within the terminal width |
| `justify` | Words are spread evenly so the output spans the full terminal width |

Any other value (e.g. `--align=diagonal`, `--align=`, `--Align=center`) will print the usage message and exit.

---

## Banner Fonts

Three banner fonts are available:

| Font | Description |
|---|---|
| `standard` | Classic bold ASCII art style |
| `shadow` | Lightweight shadow-style characters |
| `thinkertoy` | Playful toy-block style characters |

If no banner is specified, `standard` is used by default.

---

## Examples

**Center aligned:**
```bash
go run . --align=center "hello" standard
```
```
        _          _   _          
       | |        | | | |         
       | |__   ___| | | | ___     
       |  _ \ / _ \ | | |/ _ \    
       | | | |  __/ | | | (_) |   
       |_| |_|\___/_| |_|\___/    
```

**Right aligned:**
```bash
go run . --align=right "hello" shadow
```
```
                                          _| _| _|
                                    _|_|_| _|_|   
                                    _| _| _|_|_|_|
                                    _| _| _|      
                                    _| _| _|_|_|  
```

**Left aligned:**
```bash
go run . --align=left "Hello There" standard
```
```
 _   _      _ _         _____
| | | |    | | |       |_   _|
| |_| | ___| | | ___     | |  
|  _  |/ _ \ | |/ _ \    | |  
| | | |  __/ | | (_) |  _| |_ 
|_| |_|\___|_|_|\___/  |_____|
```

**Justify aligned:**
```bash
go run . --align=justify "how are you" shadow
```
Words are spread across the full terminal width with equal spacing between them.

**Multi-line input using `\n`:**
```bash
go run . "hello\nworld"
```
Renders `hello` and `world` as separate ASCII art blocks stacked vertically.

---

## How It Works

The program follows a clean pipeline from input to output:

**1. Argument Parsing (`main.go`)**
Reads `os.Args`, separates the optional `--align=` flag from positional arguments (string and banner name), validates all inputs, and exits with a usage message on any error.

**2. Banner Loading (`banner/loader.go`)**
Opens the chosen banner `.txt` file and reads it into a `map[rune][]string`. Each ASCII character from space (32) to tilde (126) maps to exactly 8 lines of ASCII art, using the formula `(charCode - 32) * 9 + 1` to locate each character block in the file.

**3. Rendering (`renderer/render.go`)**
Splits the input on literal `\n` sequences to support multi-line output. For each segment, loops through all 8 rows and concatenates each character's row from the banner map to build one complete line of ASCII art. Returns a `[]string` of individual lines.

**4. Terminal Width Detection (`terminal.go`)**
Uses a `syscall.Syscall` with `TIOCGWINSZ` to query the OS for the current terminal dimensions. Returns the column count as an `int`. This is called fresh on every run so the output adapts to the current terminal size.

**5. Alignment (`align.go`)**
Applies the chosen alignment to each rendered line:
- **Left:** lines returned as-is
- **Right:** prepends `termWidth - len(line)` spaces
- **Center:** prepends `(termWidth - len(line)) / 2` spaces
- **Justify:** renders each word separately, calculates equal spacing between words to fill the terminal width, and joins each of the 8 rows across all words

**6. Output (`main.go`)**
Loops through the aligned lines and prints each one with `fmt.Println`.

---

## Running Tests

Run all tests from the project root:

```bash
go test ./...
```

Run tests for a specific package:

```bash
go test ./renderer/...
go test .
```

Run with verbose output to see individual test results:

```bash
go test ./... -v
```

The test suite covers:
- Rendering normal input, empty input, and `\n`-separated input
- All four alignment types with known inputs and terminal widths
- Edge cases such as single-word justify (falls back to center)

---

## Allowed Packages

This project uses only Go standard library packages:

| Package | Used In | Purpose |
|---|---|---|
| `fmt` | main.go, banner/loader.go | Printing output and error messages |
| `os` | main.go, banner/loader.go | Reading files, accessing CLI args, stderr |
| `strings` | renderer/render.go, align.go, main.go | String splitting, trimming, padding |
| `bufio` | banner/loader.go | Efficient line-by-line file reading |
| `syscall` | terminal.go | Querying terminal dimensions from the OS |
| `unsafe` | terminal.go | Passing struct pointer to syscall |
| `testing` | *_test.go files | Unit testing |

No third-party packages are used.

---

## Edge Cases

| Scenario | Behaviour |
|---|---|
| Empty string `""` | Prints a blank line and exits cleanly |
| `\n` in input | Splits into multiple ASCII art blocks printed vertically |
| `\n\n` in input | Inserts a blank line between ASCII art blocks |
| Single word with `--align=justify` | Falls back to center alignment |
| `--align=` with no type | Prints usage message and exits |
| `--align=InvalidType` | Prints usage message and exits |
| Flag with wrong format (`-align`, `--Align`) | Prints usage message and exits |
| No arguments | Prints usage message and exits |
| More than 4 arguments | Prints usage message and exits |
| Terminal resized | Output adapts to new width on next run |