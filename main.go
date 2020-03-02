package main

import (
	word "Simple_English_wordBook/model"
	"Simple_English_wordBook/parse"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Home Handling
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

	port := GetPort()
	log.Println("[-] Listening on...", port)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

// for push on heroku
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
