package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		log.Println("テンプレートエラー:", err)
		http.Error(w, "テンプレートのレンダリングに失敗しました", http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title      string
		Background string
	}{
		Title:      "SHUEI's PORTFOLIO",
		Background: "static/image/Thelaat.png",
	}
	renderTemplate(w, "index", data)
}

func aboutMeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about-me", struct{ Title string }{"About Me"})
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "history", struct{ Title string }{"History"})
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "products", struct{ Title string }{"Product"})
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact", struct{ Title string }{"Contact"})
}
func skillsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "skills", struct{ Title string }{"Skills"})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "main", struct{ Title string }{"main"})
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/about-me", aboutMeHandler)
	http.HandleFunc("/history", historyHandler)
	http.HandleFunc("/skills", skillsHandler)
	http.HandleFunc("/products", productHandler)
	http.HandleFunc("/contact", contactHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("サーバーを起動中... http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("サーバーエラー:", err)
	}
}
