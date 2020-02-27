package main

import (
	word "Simple_English_wordBook/model"
	"Simple_English_wordBook/parse"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var m word.Word

func handleHome(w http.ResponseWriter, r *http.Request) {
	indexPage, _ := template.ParseFiles("index.html")

	data := struct {
		originWord string
		mean       string
		ex         string
	}{
		originWord: m.GetSWord(),
		mean:       m.GetMeanings(),
		ex:         m.GetExample(),
	}

	indexPage.Execute(w, data)
}

func handleParse(w http.ResponseWriter, r *http.Request) {
	defer http.Redirect(w, r, "/", http.StatusFound)
	word := strings.ToLower(parse.CleanString(r.FormValue("word")))
	m = parse.Parse(word)
}

func main() {
	fmt.Println("Server connecting...")
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/word", handleParse)
	http.ListenAndServe(":8080", nil)
}
