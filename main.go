package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Set up routes
	mux.HandleFunc("/message", messageHandler)
	mux.HandleFunc("/removemessage", removeMessageHandler)
	mux.HandleFunc("/", homeHandler)

	// Start server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse templates
	tmpl, err := template.ParseFiles(
		filepath.Join("templates", "layout.html"),
		filepath.Join("templates", "home.html"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute template
	data := struct {
		Title string
	}{
		Title: "Working Example",
	}
	err = tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	// Serve a small partial (not the whole layout)
	// tmpl, err := template.ParseFiles("templates/message.html")
	tmpl, err := template.New("message.html").ParseFiles("templates/message.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Message string
	}{
		Message: "Hello from HTMX!",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func removeMessageHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<div id='removed'></div>"))
}
