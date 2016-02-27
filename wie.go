package main

import (
    "fmt"
    "errors"
    "strings"

    "github.com/sbani/wie/engines"

    "github.com/PuerkitoBio/goquery"
    "gopkg.in/alecthomas/kingpin.v2"
)

const WieVersion = "0.0.1"
const Site = "stackoverflow.com";

func getAnswer(stackUrl string) (answer string, hasAnswer bool) {
    doc, err := goquery.NewDocument(stackUrl)
    if err != nil {
        return
    }

    first := doc.Find(".answer").First()
    if pre := first.Find("pre").First(); pre.Length() > 0 {
        answer = pre.Text()
        hasAnswer = true
    } else if code := first.Find("code").First(); code.Length() > 0 {
        answer = code.Text()
        hasAnswer = true
    }

    strings.TrimSpace(answer)

    return
}

// Start the howto query search
func howto(query string, engine engines.SearchEngine) (string, error) {
    // Get links from search engine
    links, err := engine.GetLinks(query, Site)
    if err != nil {
        return "", err
    } else if len(links) == 0 {
        return "", errors.New("No links found")
    }

    // Crawl links until first answer
    answer := ""
    for _, link := range links {
        answ, hasAnswer := getAnswer(link)
        if hasAnswer {
            answer = answ
            break
        }
    }

    return answer, nil
}

var (
  question = kingpin.Arg("question", "The question you ask").Required().String()
)

func main() {
    kingpin.Version(WieVersion)
    kingpin.Parse()

    answer, err := howto(*question, engines.CreateGoogle(false))
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(answer)
}
