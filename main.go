package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const word string = "hello"

type str struct {
	tilte string
}

func main() {
	// http.Handle("/", http.FileServer(http.Dir("")))
	// http.ListenAndServe(":5000", nil)

	var baseURL = "https://dic.daum.net/search.do?q=hello&dic=eng&search_first=Y"
	// var baseURL = "https://golang.org/"
	res, err := http.Get(baseURL)
	checkErr(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// searchCards := doc.Find(".list_search").First().Text()
	// fmt.Print(doc.Find(".list_search").First().Length())
	// s := strings.Join(strings.Fields(strings.TrimSpace(searchCards)), " ")
	// searchCards.Each(func(i int, card *goquery.Selection) {
	// 	fmt.Println(card.Find(".cleanword_type").Text())
	// })

	searchList := doc.Find(".box_example").First().Text()
	s := strings.Join(strings.Fields(strings.TrimSpace(searchList)), " ")
	fmt.Println(s)

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
