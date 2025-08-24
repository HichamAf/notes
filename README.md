```
 _   _       _             _______             _
| \ | | ___ | |_ ___  ___ |__   __| ___   ___ | |
|  \| |/ _ \| __/ _ \/ __|   | |   / _  \/ _ \| |
| |\  | (_) | ||  __/\__ \   | |  | (_) | (_) |_|
|_| \_|\___/ \__\___||___/   |_|   \___/ \___/(_)
```

## What is this?

A simple command-line tool to manage short, single-line notes.<br>
Each collection is just a plain text file – no databases, no magic.<br>
Useful for quick ideas, TODOs, or training notes.<br>

## Features
- Create or load a collection of notes (each collection = one text file)
- Show notes with numbering (001, 002, …)
- Add a new single-line note
- Delete a note by its number
- Notes persist between runs
- Input validation: ignores empty or whitespace-only notes
- Delete by index with safe cancel option (0 = cancel)<br>
- Password protection
- Notes encryption by ROT13
- Optional timestamp: each note can include a creation date/time (added automatically when enabled)
- Optional note title: notes can have an additional short title field besides the main text
- Optional tags: notes can be labeled with one or more tags for easier categorization and search
- Arrow key navigation in menu (w/s/j/k + Enter)
- Improved error handling in file operations

## Install & Run

Clone the repository and build or run:
```
git clone https://gitea.kood.tech/hichamafilali/notestool.git
cd notestool
go mod init notestool
```
Option A: build a binary
```
go build -o notestool
```

Option B: run directly
```
go run .
```

Run the tool with a collection name:
```
./notestool coding_ideas
```

If you run it with no arguments, more than one argument, or with the word help, you’ll see a help message:
```
./notestool
Usage: ./notestool [COLLECTION_NAME]
```

Starting the app with a <collection name> greets the user and shows a menu:
```
Welcome to the notes tool!

Select operation:
1. Show notes
2. Add a note
3. Delete a note
4. Exit
```

Example (show):
```
Notes:
001 - note one
002 - note two
```

Adding a note:
```
Enter the note text:
note three
```

Deleting a note:
```
Enter the number of note to remove or 0 to cancel:
3
```
After each operation, the menu shows again until you choose Exit.


## Data Storage
- Each collection is a plain text file named after the collection argument, stored in the working directory.
- Example: running ./notestool coding_ideas uses a file named coding_ideas.
- Each note occupies one line in that file.
- If the file doesn’t exist, it’s created the first time you run the tool with that collection name.


## Input Validation & Errors
- If the number of arguments is not exactly one, or the argument is help, a help message is shown and the program exits.
- The menu keeps prompting until you enter a valid option (1–4).
- Empty or whitespace-only notes are rejected when adding.
- Delete asks for a valid note number, checks bounds, and allows cancel with 0.

## Team & Contributions
- Hicham Afilali (Lead)
Implemented core features: Password protection, Encryption, Timestamp, and Titles. Initial project setup and skeleton (main.go orchestration, menu wiring, file loading/saving, argument validation). Drafted and updated README structure and features list. (Some of these implementations were later refined and fixed by the team.)
- Krishna Adhikari (Member)
Note operations: add, delete, show logic and user prompts; testing and edge cases.
- Gleb Simanov (Member)  
 Implemented tags for notes, arrow key menu navigation, improved encryption handling (ROT13 fix), robust error handling in file operations, and UX refinements (help messages, cancel deletion, whitespace trimming). Updated and expanded README with full documentation.

Everyone reviewed each other’s code. Final structure and behavior were discussed and agreed in group chat and a short sync call.

## Quick Test Checklist (Manual)


| Step | Action                                  | Expected Result                         |
|------|------------------------------------------|-----------------------------------------|
| 1    | Run with no args                        | Shows help and exits                    |
| 2    | Run with help                           | Shows help and exits                    |
| 3    | Run with a new collection               | File is created                         |
| 4    | Add multiple notes, exit, re-run        | Notes persist                           |
| 5    | Show notes                              | Numbers shown as 001, 002… in order     |
| 6    | Delete a note by number                 | Correct note is removed                 |
| 7    | Switch to a different collection        | Files are independent                   |
| 8    | Try invalid menu inputs (abc, 9, empty) | Tool reprompts gracefully               |
| 9    | Try deleting out-of-range index         | Error shown, then reprompt              |
| 10   | Try adding an empty note                | Rejected, reprompted                    |
