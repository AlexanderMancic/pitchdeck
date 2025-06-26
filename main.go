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
	// Erstellt einen neuen HTTP-ServeMux (Router)
	mux := http.NewServeMux()

	// Statische Dateien (CSS, Bilder etc.) aus dem "static" Ordner servieren
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routen für die verschiedenen Seiten und Funktionen definieren
	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/pitchdeck", pitchdeckHandler)
	mux.HandleFunc("/team", teamHandler)
	mux.HandleFunc("/nextslide", nextSlideHandler)                       // Nächste Folie
	mux.HandleFunc("/previousslide", previousSlideHandler)               // Vorherige Folie
	mux.HandleFunc("/guestbook", guestBookHandler)                       // Gästebuch
	mux.HandleFunc("/newcomment", newCommentHandler)                     // Neuer Kommentar Formular
	mux.HandleFunc("/createcomment", createCommentHandler)               // Kommentar erstellen
	mux.HandleFunc("/newguestbookentry", newGuestBookEntryHandler)       // Neuer Gästebucheintrag Formular
	mux.HandleFunc("/createguestbookentry", createGuestBookEntryHandler) // Gästebucheintrag erstellen
	mux.HandleFunc("/", rootHandler)                                     // Root-Handler

	// Server starten
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// Root-Handler für die Basis-URL
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Template-Dateien parsen
	tmpl, err := template.ParseFiles(
		filepath.Join("templates", "layout.html"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Template ausführen
	err = tmpl.ExecuteTemplate(w, "layout.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler für die Home-Seite
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

// Globale Variable für die aktuelle Foliennummer
var slideCounter int = 1

// Struct für Kommentare
type Comment struct {
	SlideNumber   int    // Foliennummer
	Name, Comment string // Name und Text des Kommentars
}

// Slice mit vordefinierten Beispiel-Kommentaren
var comments = []Comment{
	// Folie 1
	{SlideNumber: 1, Name: "Alex", Comment: "Toller Einstieg in die Präsentation!"},
	{SlideNumber: 1, Name: "Julia", Comment: "Das Problem ist gut dargestellt"},
	{SlideNumber: 1, Name: "Max", Comment: "Könnten Sie die Zahlen genauer erklären?"},

	// Folie 2
	{SlideNumber: 2, Name: "Sarah", Comment: "Die Marktanalyse ist sehr detailliert"},
	{SlideNumber: 2, Name: "Tom", Comment: "Welche Quellen haben Sie verwendet?"},
	{SlideNumber: 2, Name: "Anna", Comment: "Interessante Wachstumsprognosen"},

	// Folie 3
	{SlideNumber: 3, Name: "David", Comment: "Das Produktdesign gefällt mir"},
	{SlideNumber: 3, Name: "Lisa", Comment: "Gibt es schon einen Prototypen?"},
	{SlideNumber: 3, Name: "Paul", Comment: "Die USP sind klar herausgearbeitet"},

	// Folie 4
	{SlideNumber: 4, Name: "Emma", Comment: "Das Geschäftsmodell ist schlüssig"},
	{SlideNumber: 4, Name: "Felix", Comment: "Wann startet die Monetarisierung?"},
	{SlideNumber: 4, Name: "Hannah", Comment: "Die Preisstrategie ist nachvollziehbar"},

	// Folie 5-12
	{SlideNumber: 5, Name: "Oliver", Comment: "Das Team hat viel Erfahrung"},
	{SlideNumber: 6, Name: "Sophie", Comment: "Die Roadmap ist ambitioniert"},
	{SlideNumber: 7, Name: "Leon", Comment: "Starke Konkurrenzanalyse"},
	{SlideNumber: 8, Name: "Mia", Comment: "Die Finanzprojektionen sind konservativ"},
	{SlideNumber: 9, Name: "Ben", Comment: "Guter Überblick über die Risiken"},
	{SlideNumber: 10, Name: "Lena", Comment: "Der Exit-Plan ist realistisch"},
	{SlideNumber: 11, Name: "Jonas", Comment: "Beeindruckende Traction"},
	{SlideNumber: 12, Name: "Laura", Comment: "Überzeugender Abschluss!"},
}

// Handler für die Pitchdeck-Seite
func pitchdeckHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("pitchdeck.html").ParseFiles("templates/pitchdeck.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Daten für das Template
	data := struct {
		CurrentSlideNumber int       // Aktuelle Foliennummer
		Comments           []Comment // Kommentare für die aktuelle Folie
	}{
		CurrentSlideNumber: slideCounter,
		Comments:           comments,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler für die Team-Seite
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

// Struct für Gästebucheinträge
type GuestBookEntry struct {
	Name, Entry string // Name und Eintragstext
}

// Slice mit vordefinierten Gästebucheinträgen
var guestbook = []GuestBookEntry{
	{Name: "Alex Schmidt", Entry: "Tolle Präsentation! Wann geht's los mit der Umsetzung?"},
	{Name: "Julia Weber", Entry: "Sehr professioneller Pitch. Viel Erfolg!"},
	{Name: "Max Müller", Entry: "Interessantes Konzept. Melden Sie sich bei Interesse an einer Kooperation"},
	{Name: "Sarah Meyer", Entry: "Habe noch Fragen zur Technologie. Können wir uns austauschen?"},
	{Name: "Tom Wagner", Entry: "Beeindruckende Zahlen. Wann ist die nächste Funding-Runde?"},
	{Name: "Anna Fischer", Entry: "Das Team hat mich überzeugt. Bin gespannt auf die Entwicklung"},
	{Name: "David Becker", Entry: "Gut strukturierte Präsentation. Die Slides waren sehr klar"},
	{Name: "Lisa Hoffmann", Entry: "Würde gerne mehr über die Kundensegmentierung erfahren"},
	{Name: "Paul Schulz", Entry: "Top Pitches heute, aber Ihre Präsentation war die beste!"},
	{Name: "Emma Köhler", Entry: "Haben Sie schon mit potenziellen Partnern gesprochen?"},
}

// Handler für das Gästebuch
func guestBookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("guestbook.html").ParseFiles("templates/guestbook.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Daten für das Template
	data := struct {
		GuestBookEntries []GuestBookEntry // Alle Gästebucheinträge
	}{
		GuestBookEntries: guestbook,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler für das Formular zum Erstellen eines neuen Gästebucheintrags
func newGuestBookEntryHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("newguestbookentry.html").ParseFiles("templates/newguestbookentry.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler zum Speichern eines neuen Gästebucheintrags
func createGuestBookEntryHandler(w http.ResponseWriter, r *http.Request) {
	// Formulardaten parsen
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Formularfelder auslesen
	name := r.FormValue("name")
	entry := r.FormValue("entry")

	// Neuen Eintrag hinzufügen
	guestbook = append(guestbook, GuestBookEntry{
		Name:  name,
		Entry: entry,
	})

	// Gästebuch-Seite neu anzeigen
	guestBookHandler(w, r)
}

// Handler für das Formular zum Erstellen eines neuen Kommentars
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

// Handler zum Speichern eines neuen Kommentars
func createCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Formulardaten parsen
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Formularfelder auslesen
	name := r.FormValue("name")
	comment := r.FormValue("comment")

	// Neuen Kommentar hinzufügen
	comments = append(comments, Comment{
		SlideNumber: slideCounter, // Aktuelle Foliennummer
		Name:        name,
		Comment:     comment,
	})

	// Pitchdeck-Seite neu anzeigen
	pitchdeckHandler(w, r)
}

// Handler für die nächste Folie
func nextSlideHandler(w http.ResponseWriter, r *http.Request) {
	// Folienzähler erhöhen (max. 12 Folien)
	if slideCounter < 12 {
		slideCounter++
	}

	// HTML für die neue Folie und Kommentare generieren
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

// Handler für die vorherige Folie
func previousSlideHandler(w http.ResponseWriter, r *http.Request) {
	// Folienzähler verringern (min. 1)
	if slideCounter > 1 {
		slideCounter--
	}

	// HTML für die neue Folie und Kommentare generieren
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

// Hilfsfunktion zum Rendern der Kommentare als HTML
func renderCommentsHTML(comments []Comment) string {
	var sb strings.Builder
	sb.WriteString(`<div id="comments" hx-swap-oob="true">`)

	// Nur Kommentare für die aktuelle Folie anzeigen
	for _, c := range comments {
		if c.SlideNumber == slideCounter {
			sb.WriteString(fmt.Sprintf("<h3>%s</h3><div>%s</div>", c.Name, c.Comment))
		}
	}

	sb.WriteString("</div>")
	return sb.String()
}
