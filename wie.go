package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const site = "stackoverflow.com"

const googleHost = "https://www.google.com"
const googleURI = "/search?q=site:%s+%s"

// answer holds all information related to an answer
type answer struct {
	url                 string
	html                *goquery.Selection
	shortText, longText string
	votes               int
}

// newAnswer creates a nem answer
func newAnswer(url string) (*answer, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	first := doc.Find(".answer").First()
	if first.Length() == 0 {
		return nil, errors.New("No answer found")
	}

	aw := &answer{}
	aw.url = url
	aw.html = first

	return aw, nil
}

// parseForCode parses the html and tries to find a code or a pre element. Just the first one
func (a *answer) parseForCode() {
	if pre := a.html.Find("pre").First(); pre.Length() > 0 {
		a.shortText = strings.TrimSpace(pre.Text())
	} else if code := a.html.Find("code").First(); code.Length() > 0 {
		a.shortText = strings.TrimSpace(code.Text())
	}
}

// parseForText parses the answer html for the complete text from .post-text
func (a *answer) parseForText() {
	if text := a.html.Find(".post-text").First(); text.Length() > 0 {
		a.longText = strings.TrimSpace(text.Text())
	}
}

// parseForVotes parses the answer html for the votes (up/down votes)
func (a *answer) parseForVotes() {
	votes, _ := strconv.Atoi(a.html.Find(".vote-count-post").First().Text())
	a.votes = votes
}

// createGoogleSearchUrl is used to create the google search URL.
// Example: `https://www.google.com/search?q=site:stackoverflow.com+windows+get+date+command+line`
func createGoogleSearchURL(query, site string) string {
	uri := fmt.Sprintf(googleURI, site, url.QueryEscape(query))
	return googleHost + uri
}

// getLinks for a `query` and `site`.
// `site`: What site you want to get links for
// `query`: Search query you want to search on `site` for.
func getLinks(query, site string) (links []string, err error) {
	doc, err := goquery.NewDocument(createGoogleSearchURL(query, site))
	if err != nil {
		return
	}

	filterLink := func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			links = append(links, googleHost+href)
		}
	}

	doc.Find("a.l").Each(filterLink)
	if len(links) == 0 {
		doc.Find(".r a").Each(filterLink)
	}

	return
}

// search with specific query and return error or answer
func search(query string) (*answer, error) {
	// Get links from search engine
	links, err := getLinks(query, site)
	if err != nil {
		return nil, err
	} else if len(links) == 0 {
		return nil, errors.New("No links found")
	}

	// Crawl links until first answer
	for _, link := range links {
		aw, err := newAnswer(link)
		if err == nil {
			aw.parseForCode()
			go aw.parseForText()
			go aw.parseForVotes()

			return aw, nil
		}
	}

	return nil, errors.New("No answer found")
}

func main() {
	query := strings.Join(os.Args[1:], " ")

	if query == "" {
		fmt.Println("Missing query")
		os.Exit(1)
		return
	}

	fmt.Println("Running...")
	aw, err := search(query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	fmt.Println(aw.shortText)
	fmt.Println("Legend: [a] Show all, [Enter|q] Quit")
	fmt.Print("Enter text: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(strings.ToLower(text))

	switch text {
	case "":
	case "q":
		return
	case "a":
		fmt.Println(aw.longText)
	}
}
