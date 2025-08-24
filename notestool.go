package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Load notes from file; if file doesn't exist -> empty list
func LoadNotes(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{} // no file yet -> empty
		}
		fmt.Println(Red+"Error opening file:"+Reset, err)
		return []string{}
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		notes = append(notes, rot13(line)) // decrypt when loading
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(Red+"Error reading file:"+Reset, err)
	}
	return notes
}

// Save notes back to file
func SaveNotes(filename string, notes []string) {
	file, err := os.Create(filename) // overwrite file
	if err != nil {
		fmt.Println(Red+"Error saving notes:"+Reset, err)
		return
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Println(Red+"Error closing file:"+Reset, cerr)
		}
	}()
	writer := bufio.NewWriter(file)
	for _, note := range notes {
		if _, err := writer.WriteString(rot13(note) + "\n"); err != nil {
			fmt.Println(Red+"Error writing file:"+Reset, err)
			return
		}
	}
	if err := writer.Flush(); err != nil {
		fmt.Println(Red+"Error flushing file:"+Reset, err)
	}
}

// Show notes with numbering
func ShowNotes(collection string, notes []string) {
	fmt.Println(Blue + "\n+---------------------------------------------+" + Reset)
	fmt.Printf(Blue +  "| COLLECTION: " + Red + "%s\n" + Reset, collection)
	//fmt.Println(Blue + "|--------------------------------------------|" + Reset)
	
	//fmt.Println("| ", collection)

	if len(notes) == 0 {
		//fmt.Println(Blue + "|--------------------------------------------|" + Reset)
		fmt.Println(Blue +  "|" + Red + " No notes yet.                              " + Blue + "|" + Reset)
		fmt.Println(Blue + "+--------------------------------------------+" + Reset)
		return
	}
	fmt.Println(Blue + "+------+------+------------+------------------+" + Reset)
	fmt.Println(Blue + "|  ID  | Tag  |    Date    |       Note       |" + Reset)
	fmt.Println(Blue + "+------+------+------------+------------------+" + Reset)
	//fmt.Println(Blue + "|                   NOTES                    |" + Reset)
	//fmt.Println(Blue + "|--------------------------------------------|" + Reset)
	for i, note := range notes {
		fmt.Printf(Blue + "| %03d  | %s" + Reset + "\n", i+1, note)
	}
	fmt.Println(Blue +"+---------------------------------------------+" + Reset)
}

// Add note and save (asks for optional date string from user)
func AddNote(filename string, notes []string, newNote string) []string {
	// ignore empty/whitespace-only note
	if strings.TrimSpace(newNote) == "" {
		fmt.Println(Red + "Empty note, not added." + Reset)
		return notes
	}

	// ask user for optional date string (since time package is not allowed)
	var date string
	for {
		fmt.Println(Yellow + "Enter date (optional, e.g. 18.08.2025). Leave empty to skip:" + Reset)
		date = ReadInput()

		if date == "" {
			// user skipped date → keep fixed-length blank
        	date = "          " // 10 spaces, same width as dd.mm.yyy
			break // user skipped date
		}

		parts := strings.Split(date, ".")
		if len(parts) != 3 {
			fmt.Println(Red + "Invalid format! Use dd.mm.yyyy" + Reset)
			continue
		}

		day, dErr := strconv.Atoi(parts[0])
		month, mErr := strconv.Atoi(parts[1])
		year, yErr := strconv.Atoi(parts[2])

		if dErr != nil || mErr != nil || yErr != nil ||
			day < 1 || day > 31 ||
			month < 1 || month > 12 ||
			year < 1 {
			fmt.Println(Red + "Invalid values! Use dd.mm.yyyy" + Reset)
			continue
		}

		break // valid date
	}

	// build final note text (prefix date if provided)
	final := newNote
	if strings.TrimSpace(date) != "" {
		final = "| " + strings.TrimSpace(date) + " | " + newNote
	}

	// ask user for optional tags (comma-separated), e.g. "work,idea"
	fmt.Println(Yellow + "Enter tags (optional, comma-separated). Leave empty to skip:" + Reset)
	tagsInput := ReadInput()
	tagsInput = strings.TrimSpace(tagsInput)

	// if tags provided, normalize and prepend to the note
	if tagsInput != "" {
		// make commas spaced a bit nicer: "a,b" -> "a, b"
		tagsInput = strings.ReplaceAll(tagsInput, ",", ", ")
		final = "  " + tagsInput + "  " + final
	}

	notes = append(notes, final)
	SaveNotes(filename, notes)
	return notes
}

// Delete note by number and save
func DeleteNote(filename string, notes []string, index int) []string {
	// index is 1-based in the UI; 0 cancels
	if index == 0 {
		fmt.Println(Red + "Canceled." + Reset)
		return notes
	}
	if index < 1 || index > len(notes) {
		fmt.Println(Red + "Invalid note number." + Reset)
		return notes
	}

	i := index - 1
	// remove element at i
	notes = append(notes[:i], notes[i+1:]...)

	// persist changes
	SaveNotes(filename, notes)
	return notes
}

// encryption
func rot13(s string) string {
	var b strings.Builder
	for _, c := range s {
		switch {
		case c >= 'A' && c <= 'Z':
			b.WriteRune((c-'A'+13)%26 + 'A')
		case c >= 'a' && c <= 'z':
			b.WriteRune((c-'a'+13)%26 + 'a')
		default:
			b.WriteRune(c)
		}
	}
	return b.String()
}

// MenuSelect shows a highlightable menu. Use: w/k = up, s/j = down, Enter = choose, or type 1..n.
func MenuSelect(options []string) int {
	idx := 0
	for {
		ClearScreen()
		fmt.Println(Blue +"+--------------------------------------------+" + Reset)
		fmt.Println(Blue +"|" + Yellow + Bold + "        📒 Welcome to Notes Tool 📒" + Reset + Blue + "         |" + Reset)
		fmt.Println(Blue +"+--------------------------------------------+" + Reset)
		fmt.Println(Blue +"| MENU: Select operation                     |" + Reset)
		fmt.Println(Blue +"+--------------------------------------------+" + Reset)
		//fmt.Println(Blue + "|    " + Green + " Select operation:                      " + Blue + "|" + Reset)
		//fmt.Println(Blue + "|--------------------------------------------|" + Reset)
		for i, opt := range options {
			marker := Blue + "|   " + Reset
			if i == idx {
				marker = Blue + "|" + Yellow + " > " + Reset
			}
			fmt.Println(marker + opt)
		}
		fmt.Println(Blue +"+--------------------------------------------+" + Reset)
		fmt.Printf(Yellow + "\nUse w/k (up/down), Enter (choose) or type number (1-%d): " + Reset, len(options))

		input := ReadInput() // line-based read from stdin
		in := strings.TrimSpace(strings.ToLower(input))

		switch in {
		case "": // Enter → choose current
			return idx
		case "w":
			if idx > 0 {
				idx--
			} else {
				idx = len(options) - 1
			}
		case "s":
			if idx < len(options)-1 {
				idx++
			} else {
				idx = 0
			}
		default:
			// quick-select by number
			n, err := strconv.Atoi(in)
			if err == nil && n >= 1 && n <= len(options) {
				return n - 1
			}
			// invalid input → redraw and continue
		}
	}
}
