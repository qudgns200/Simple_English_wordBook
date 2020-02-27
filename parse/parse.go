package parse

import (
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

// method for type searchWord
func (this searchWord) String() string {
	return this.sWord + " " + this.meanings + " " + this.example
}

func Parse(word string) *searchWord {
	var result searchWord
	returnValue := getSearchResult(word, result)
	return returnValue
}

// Get the result after searching
func getSearchResult(word string, result searchWord) *searchWord {
	var baseURL = "https://dic.daum.net/search.do?q=" + word + "&dic=eng&search_first=Y"
	res, err := http.Get(baseURL)
	checkCode(res)
	checkErr(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	meanings := doc.Find(".list_search").First().Text()
	example := doc.Find(".box_example").First().Text()

	returnValue := setSearchWord(word, meanings, example, result)

	return returnValue

}

// Set value of searchWord struct
func setSearchWord(word string, meanings string, example string, result searchWord) *searchWord {
	trimMeanings := CleanString(meanings)
	trimExample := CleanString(example)

	result.sWord = word
	result.meanings = trimMeanings
	result.example = trimExample

	return &result
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
