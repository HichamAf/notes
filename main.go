package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ANSI color codes for styling terminal output
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
	Reset  = "\033[0m"
)

func main() {
	// check if exactly one argument is passed (the collection name)
	if len(os.Args) != 2 || os.Args[1] == "help" {
		// fmt.Println("Usage: ./notestool [COLLECTION_NAME]")
		PrintHelp()
		return
	}

	if !CheckPassword() {
		return // stop program if wrong password
	}

	collection := os.Args[1]

	// Load or create the notes file
	notes := LoadNotes(collection)

	// Main loop (keeps running until user chooses Exit)
	for {
		// MenuSelect will render its own selectable list and handle w/k/s/j or numbers.
		options := []string{
			Blue + "|" + Green + " 1. Show notes                          " + Blue + "|" + Reset,
			Blue + "|" + Green + " 2. Add a note                          " + Blue + "|" + Reset,
			Blue + "|" + Green + " 3. Delete a note                       " + Blue + "|" + Reset,
			Blue + "|" + Green + " 4. Exit                                " + Blue + "|" + Reset,
		}

		selected := MenuSelect(options)

		switch selected {
		case 0: // Show notes
			ShowNotes(collection, notes)
			Pause()

		case 1: // Add note
			fmt.Println(Yellow + "\nEnter the note text:" + Reset)
			note := ReadInput()
			notes = AddNote(collection, notes, note)
			Pause()

		case 2: // Delete note
			ShowNotes(collection, notes) // display current notes
			fmt.Println(Yellow + "\nEnter the number of note to remove or 0 to cancel:" + Reset)
			numStr := ReadInput()
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println(Red + "Invalid input, must be a number" + Reset)
				Pause()
				continue
			}
			if num == 0 {
				continue // cancel deletion
			}
			notes = DeleteNote(collection, notes, num)
			Pause()

		case 3: // Exit
			return
		}
	}
}

// helper function to read input from user
func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text) // remove newline
}

// PrintHelp shows how to use the notes tool
func PrintHelp() {
	fmt.Println("\nUsage: ./notestool [COLLECTION_NAME]")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ./notestool coding_ideas   # manage a collection called 'coding_ideas'")
	fmt.Println("  ./notestool work_notes     # manage a collection called 'work_notes'")
	fmt.Println()
	fmt.Println("If the collection does not exist, it will be created automatically.")
	fmt.Println()
}

// ClearScreen clears the terminal
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Pause waits for the user to press Enter
func Pause() {
	fmt.Println(Yellow + "\nPress ENTER to continue..." + Reset)
	ReadInput()
}

// simple password (hardcoded)
const PASSWORD = ""

func CheckPassword() bool {
	fmt.Print(Yellow + "\nEnter password: " + Reset)
	reader := bufio.NewReader(os.Stdin)
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimSpace(pass)

	if pass == PASSWORD {
		return true
	}
	fmt.Println(Red + "Wrong password. Access denied." + Reset)
	return false
}
