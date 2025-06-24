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
	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/pitchdeck", pitchdeckHandler)
	mux.HandleFunc("/team", teamHandler)
	mux.HandleFunc("/nextslide", nextSlideHandler)
	mux.HandleFunc("/previousslide", previousSlideHandler)
	mux.HandleFunc("/", rootHandler)

	// Start server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Parse templates
	tmpl, err := template.ParseFiles(
		filepath.Join("templates", "layout.html"),
		// filepath.Join("templates", "home.html"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute template
	err = tmpl.ExecuteTemplate(w, "layout.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("home.html").ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var slideCounter int = 1

func pitchdeckHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("pitchdeck.html").ParseFiles("templates/pitchdeck.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func teamHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("team.html").ParseFiles("templates/team.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func nextSlideHandler(w http.ResponseWriter, r *http.Request) {

	if slideCounter < 4 {
		slideCounter++
	}

	html := fmt.Sprintf(
		`
		<img src='/static/png/Folie%v.PNG' alt='Pitchdeckfolie'>
		<span id='pageNumber' hx-swap-oob="true">%v</span>`,
		slideCounter,
		slideCounter,
	)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func previousSlideHandler(w http.ResponseWriter, r *http.Request) {

	if slideCounter > 1 {
		slideCounter--
	}

	html := fmt.Sprintf(
		`
		<img src='/static/png/Folie%v.PNG' alt='Pitchdeckfolie'>
		<span id='pageNumber' hx-swap-oob="true">%v</span>`,
		slideCounter,
		slideCounter,
	)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}
