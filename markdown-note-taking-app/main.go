package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"markdown-note-taking-app/notes"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var noteRepo notes.NoteRepository

func createNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	if title == "" || content == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing title or content")
		return
	}
	note := &notes.Note{Title: title, Content: content}
	id, err := noteRepo.Create(note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not create note")
		return
	}
	fmt.Fprintf(w, "Note created with ID %d", id)
}

func listNotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	notesList, err := noteRepo.List()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not list notes")
		return
	}
	json.NewEncoder(w).Encode(notesList)
}

func readNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	var id int64
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID")
		return
	}
	note, err := noteRepo.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Note not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	var id int64
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID")
		return
	}
	err = noteRepo.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not delete note")
		return
	}
	fmt.Fprint(w, "Note deleted")
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	idStr := r.FormValue("id")
	var id int64
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID")
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	if title == "" || content == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing title or content")
		return
	}
	note := &notes.Note{ID: id, Title: title, Content: content}
	err = noteRepo.Update(note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not update note")
		return
	}
	fmt.Fprint(w, "Note updated")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	//1. Param input for multipart file upload
	r.ParseMultipartForm(200 << 20) // Maximum of 200MB file allowed

   file, _, err := r.FormFile("file")
   title := r.FormValue("title")
   if err != nil {
	   w.WriteHeader(http.StatusBadRequest)
	   fmt.Fprint(w, "Error parsing form file")
	   return
   }

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error reading file")
		return
	}

	note := &notes.Note{Title: title, Content: string(fileBytes)}
	id, err := noteRepo.Create(note)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error saving file")
		return
	}

   defer file.Close()

   w.WriteHeader(http.StatusOK)
   fmt.Fprint(w, "File uploaded successfully and created note with ID: ", id)
}


func checkGrammarAndSpelling(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Check grammar and spelling")
}

// CORS middleware for development
func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func startFrontendServer() {
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func startBackendServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/createNote", createNote)
	mux.HandleFunc("/api/upload", UploadHandler)
	mux.HandleFunc("/api/checkGrammarAndSpelling", checkGrammarAndSpelling)
	mux.HandleFunc("/api/listNotes", listNotes)
	mux.HandleFunc("/api/readNote", readNote)
	mux.HandleFunc("/api/deleteNote", deleteNote)
	mux.HandleFunc("/api/updateNote", updateNote)
	log.Fatal(http.ListenAndServe(":8080", withCORS(mux)))
}

func main() {
	db, err := sql.Open("sqlite3", "./notes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		content TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}
	noteRepo = notes.NewSqliteNoteRepository(db)
	go startFrontendServer()
	startBackendServer()
}