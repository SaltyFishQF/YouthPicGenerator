package utils

import (
	"../model"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

//"http://news.cyol.com/node_67071.htm"

var lessons []model.Lesson
var LESSON_JSON string

func getHtml(url string) *goquery.Document {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func Spider() {
	doc := getHtml("http://news.cyol.com/node_67071.htm")
	// Find the review items
	doc.Find(".movie-list").Find("li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		link, _ := s.Find("h3").Find("a").Attr("href")
		link = link[:strings.LastIndex(link, "/")]

		innerDoc := getHtml(link)
		title := innerDoc.Find(".cont_h").Find("h1").Text()

		if title == "" {
			title = link[strings.LastIndex(link, "/")+1:]
		}

		lesson := model.Lesson{title, link}
		lessons = append(lessons, lesson)
	})
	json, _ := json.Marshal(lessons)
	LESSON_JSON = string(json)
}
