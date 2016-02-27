package engines

import (
    "fmt"
    "net/url"

    "github.com/PuerkitoBio/goquery"
)

// Google search engine
type Google struct {
    Host, Uri string
    SSL bool
}

// Create google search engine
func CreateGoogle(ssl bool) Google {
    var schema string
    if ssl {
        schema = "https"
    } else {
        schema = "http"
    }

    return Google{
        schema + "://www.google.com",
        "search?q=site:%s+%s",
        ssl,
    }
}

// Get links for specific search `query`
func (google Google) GetLinks(query, site string) (links []string, err error) {
    doc, err := goquery.NewDocument(google.createSearchUrl(query, site))
    if err != nil {
        return
    }

    filterLink := func(i int, s *goquery.Selection) {
        href, exists := s.Attr("href")
        if exists {
            links = append(links, google.Host + href)
        }
    }

    doc.Find("a.l").Each(filterLink)
    if len(links) == 0 {
        doc.Find(".r a").Each(filterLink)
    }

    return
}

// Create search url for specific site and with query.
func (google Google) createSearchUrl(query, site string) string {
    uri := fmt.Sprintf(google.Uri, site, url.QueryEscape(query))
    return google.Host + "/" + uri
}
