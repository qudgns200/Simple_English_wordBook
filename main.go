package main

import (
	word "Simple_English_wordBook/model"
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
	var arr []word.Word
	d := parse.Parse(r.FormValue("term"), 0)
	e := parse.Parse(r.FormValue("term"), 1)

	arr = append(arr, d)
	arr = append(arr, e)

	// create json response from struct
	answer3, err := json.Marshal(arr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(answer3)
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/word", ajaxHandler)
	http.ListenAndServe(":8080", nil)
}
