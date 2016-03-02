package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/sbani/wie/engines"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"gopkg.in/alecthomas/kingpin.v2"
)

const wieVersion = "0.0.2"
const site = "stackoverflow.com"

// Get the short version of the answer with `<code>` or `<pre>` only.
func getCodeFromAnswer(s *goquery.Selection) (answer string, hasAnswer bool) {
	if pre := s.Find("pre").First(); pre.Length() > 0 {
		answer = pre.Text()
		hasAnswer = true
	} else if code := s.Find("code").First(); code.Length() > 0 {
		answer = code.Text()
		hasAnswer = true
	}

	return
}

// Get the complete answer
func getCompleteAnswer(s *goquery.Selection) (answer string, hasAnswer bool) {
	if text := s.Find(".post-text").First(); text.Length() > 0 {
		answer = text.Text()
		hasAnswer = true
	}

	return
}

// Get the first answer from a given stackoverflow.com url.
func getAnswer(stackURL string, showAll bool) (votes int, answer string, hasAnswer bool) {
	doc, err := goquery.NewDocument(stackURL)
	if err != nil {
		return
	}

	first := doc.Find(".answer").First()
	if showAll {
		answer, hasAnswer = getCompleteAnswer(first)
	} else {
		answer, hasAnswer = getCodeFromAnswer(first)
	}

	answer = strings.TrimSpace(answer)

	votes, err = strconv.Atoi(first.Find(".vote-count-post").First().Text())

	return
}

// Start the howto query search
func howto(query string, engine engines.SearchEngine, showAll bool) (votes int, answer string, err error) {
	// Get links from search engine
	links, err := engine.GetLinks(query, site)
	if err != nil {
		return
	} else if len(links) == 0 {
		err = errors.New("No links found")
		return
	}

	// Crawl links until first answer
	for _, link := range links {
		vts, answ, hasAnswer := getAnswer(link, showAll)
		if hasAnswer {
			answer = answ
			votes = vts
			break
		}
	}

	return
}

var (
	question   = kingpin.Arg("question", "The question you ask").Required().String()
	printAll   = kingpin.Flag("all", "Display the hole answer").Short('a').Bool()
	printVotes = kingpin.Flag("votes", "Print answer's votes").Short('v').Bool()
)

func main() {
	kingpin.Version(wieVersion)
	kingpin.Parse()

	votes, answer, err := howto(*question, engines.CreateGoogle(false), *printAll)
	if err != nil {
		fmt.Println(err)
	}

	if *printVotes {
		color.Cyan(fmt.Sprintf("Votes: %d", votes))
	}

	fmt.Println(answer)
}
