package main

import (
	"Simple_English_wordBook/parse"
	"encoding/json"
	"html/template"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Pass template to http.ResponseWriter.
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// AJAX Request Handler
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	d := parse.Parse(r.FormValue("term"))

	// create json response from struct
	answer, err := json.Marshal(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(answer)
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/word", ajaxHandler)
	http.ListenAndServe(":8080", nil)
}
