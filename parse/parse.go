package parse

import (
	word "Simple_English_wordBook/model"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Parse(term string) word.Word {
	var result word.Word
	returnValue := getSearchResult(term, result)
	return returnValue
}

// Get the result after searching
func getSearchResult(term string, result word.Word) word.Word {
	var baseURL = "https://dic.daum.net/search.do?q=" + term + "&dic=eng&search_first=Y"
	res, err := http.Get(baseURL)
	checkCode(res)
	checkErr(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	meanings := doc.Find(".list_search").First().Text()
	example := doc.Find(".box_example").First().Text()

	returnValue := setSearchWord(term, meanings, example, result)

	return returnValue

}

// Set value of SearchWord struct
func setSearchWord(word string, meanings string, example string, result word.Word) word.Word {
	trimMeanings := CleanString(meanings)
	trimExample := CleanString(example)

	result.SWord = word
	result.Meanings = trimMeanings
	result.Example = trimExample

	return result
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
