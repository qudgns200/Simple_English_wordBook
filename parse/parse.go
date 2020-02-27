package parse

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type searchWord struct {
	sWord    string
	meanings string
	example  string
}

func Parse(word string) {
	var result searchWord
	c := make(chan searchWord)
	getSearchResult(word, c)
}

// Get the result after searching
func getSearchResult(word string, c chan<- searchWord) {
	var baseURL = "https://dic.daum.net/search.do?q=" + word + "&dic=eng&search_first=Y"
	res, err := http.Get(baseURL)
	checkCode(res)
	checkErr(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	meanings := doc.Find(".list_search").First().Text()
	example := doc.Find(".box_example").First().Text()

	setSearchWord(word, meanings, example, c)

}

// Set value of searchWord struct
func setSearchWord(word string, meanings string, example string, c chan<- searchWord) {
	trimMeanings := CleanString(meanings)
	trimExample := CleanString(example)

	c <- searchWord{
		sWord: word,
		meanings: meanings
		example: example
	}
}

// CleanString cleans a String
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

// if it occurs error, print error message
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Check the connect
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
