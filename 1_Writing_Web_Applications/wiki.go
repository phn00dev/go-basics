package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func (p *Page) loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there , I love %s!", r.URL.Path[1:])
}

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	var p Page
// 	page, _ := p.loadPage(title)
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
// }

// func editHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/edit/"):]
// 	var p Page
// 	page, err := p.loadPage(title)
// 	if err != nil {
// 		p = Page{Title: title}
// 	}
// 	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
// 		"<form action=\"/save/%s\" method=\"POST\">"+
// 		"<textarea name=\"body\">%s</textarea><br>"+
// 		"<input type=\"submit\" value=\"Save\">"+
// 		"</form>", page.Title, page.Title, page.Body)
// }

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		log.Println("error get title :", err)
		return
	}
	var page Page
	p, err := page.loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "edit", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	var page Page
	p, _ := page.loadPage(title)
	renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	if err := p.save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+p.Title, http.StatusFound)
	return
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	log.Println(m)
	if m == nil {

		http.NotFound(w, r)
		return "", errors.New("invalid Page title")
	}
	log.Println(m[2])
	return m[2], nil
}

// func makeHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		m := validPath.FindStringSubmatch(r.URL.Path)
// 		if m == nil {
// 			http.NotFound(w, r)
// 			return
// 		}
// 		fn(w, r, m[2])
// 	}
// }

func main() {

	// p1 := &Page{Title: "TestPage", Body: []byte("This is a simple Page")}
	// p1.save()
	// p2, _ := p1.loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
