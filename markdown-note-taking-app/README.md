# Markdown Note Taking App

A simple markdown note taking app written in Go.

## Functional requirements

- Create notes: uploading a markdown file, check grammar and spelling
- List notes: listing all notes with pagination
- Read notes: reading a note content rendered as HTML
- Delete notes: deleting a note
- Update notes: updating a note
- Use markdown as format for notes content, but for render notes use HTML

## Non functional requirements

- The notes needs to be persistent in a secure way, use a database, SQLite is preferred


## Tech stack

### Backend

- Go 1.22.2
- sqlite

### Frontend

To avoid using a framework, we will use HTML, CSS and JS.

- HTML
- CSS
- JS



## How to run the app


- Backend: `go run main.go`
- Frontend: `task frontend-serve` or use Live Server extension in VSCode
