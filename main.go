package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
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
	mux.HandleFunc("/guestbook", guestBookHandler)
	mux.HandleFunc("/newcomment", newCommentHandler)
	mux.HandleFunc("/createcomment", createwCommentHandler)
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

type Comment struct {
	SlideNumber   int
	Name, Comment string
}

var comments = []Comment{
	{SlideNumber: 1, Name: "Alex", Comment: "Hallo Welt"},
	{SlideNumber: 1, Name: "Jamie", Comment: "Nice intro!"},
	{SlideNumber: 2, Name: "Chris", Comment: "Can you explain this more?"},
	{SlideNumber: 2, Name: "Sam", Comment: "Interesting point here."},
	{SlideNumber: 3, Name: "Phillip", Comment: "Was geht"},
	{SlideNumber: 3, Name: "Taylor", Comment: "Great visuals!"},
	{SlideNumber: 4, Name: "Jordan", Comment: "Wrap-up was clear."},
	{SlideNumber: 4, Name: "Morgan", Comment: "Thanks for the presentation!"},
}

func pitchdeckHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("pitchdeck.html").ParseFiles("templates/pitchdeck.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		CurrentSlideNumber int
		Comments           []Comment
	}{
		CurrentSlideNumber: slideCounter,
		Comments:           comments,
	}

	err = tmpl.Execute(w, data)
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

func guestBookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("guestbook.html").ParseFiles("templates/guestbook.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func newCommentHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("newcomment.html").ParseFiles("templates/newcomment.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func createwCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Extract form fields
	name := r.FormValue("name")
	comment := r.FormValue("comment")

	comments = append(comments, Comment{
		SlideNumber: slideCounter,
		Name:        name,
		Comment:     comment,
	})

	pitchdeckHandler(w, r)
}

func nextSlideHandler(w http.ResponseWriter, r *http.Request) {

	if slideCounter < 4 {
		slideCounter++
	}

	html := fmt.Sprintf(
		`
		<img src='/static/png/Folie%v.PNG' alt='Pitchdeck Foliennummer %v'>
		<span id='pageNumber' hx-swap-oob="true">%v</span>
		%s`,
		slideCounter,
		slideCounter,
		slideCounter,
		renderCommentsHTML(comments),
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
		<img src='/static/png/Folie%v.PNG' alt='Pitchdeck Foliennummer %v'>
		<span id='pageNumber' hx-swap-oob="true">%v</span>
		%s`,
		slideCounter,
		slideCounter,
		slideCounter,
		renderCommentsHTML(comments),
	)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func renderCommentsHTML(comments []Comment) string {

	var sb strings.Builder
	sb.WriteString(`<div id="comments" hx-swap-oob="true">`)

	for _, c := range comments {
		if c.SlideNumber == slideCounter {
			sb.WriteString(fmt.Sprintf("<h3>%s</h3><div>%s</div>", c.Name, c.Comment))
		}
	}

	sb.WriteString("</div>")
	return sb.String()
}
