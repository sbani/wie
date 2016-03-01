package engines

import (
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// Google search engine
type Google struct {
	Host, URI string
	SSL       bool
}

// CreateGoogle is the func to create the search engine with default values
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

// GetLinks for a `query` and `site`.
// `site`: What site you want to get links for
// `query`: Search query you want to search on `site` for.
func (google Google) GetLinks(query, site string) (links []string, err error) {
	doc, err := goquery.NewDocument(google.createSearchURL(query, site))
	if err != nil {
		return
	}

	filterLink := func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			links = append(links, google.Host+href)
		}
	}

	doc.Find("a.l").Each(filterLink)
	if len(links) == 0 {
		doc.Find(".r a").Each(filterLink)
	}

	return
}

// createSearchURL is used to create the google search URL.
// Example: `https://www.google.de/search?q=site:stackoverflow.com+windows+get+date+command+line`
func (google Google) createSearchURL(query, site string) string {
	uri := fmt.Sprintf(google.URI, site, url.QueryEscape(query))
	return google.Host + "/" + uri
}
