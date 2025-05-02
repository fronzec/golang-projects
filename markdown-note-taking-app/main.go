package main

import (
	"fmt"
	"log"
	"net/http"
	"markdown-note-taking-app/notes"
)


func createNote(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create note")
}

func checkGrammarAndSpelling(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Check grammar and spelling")

}

func listNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "List notes")
}

func readNote(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Read note")
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete note")
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Update note")
}

func startFrontendServer() {
		fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func startBackendServer() {
	// todo change to router and use verbs GET, POST, PUT, DELETE

	mux := http.NewServeMux()
	mux.HandleFunc("/api/createNote", createNote)
	mux.HandleFunc("/api/upload", notes.UploadHandler)
	mux.HandleFunc("/api/checkGrammarAndSpelling", checkGrammarAndSpelling)
	mux.HandleFunc("/api/listNotes", listNotes)
	mux.HandleFunc("/api/readNote", readNote)
	mux.HandleFunc("/api/deleteNote", deleteNote)
	mux.HandleFunc("/api/updateNote", updateNote)
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func main() {
	go startFrontendServer()
	startBackendServer()
}