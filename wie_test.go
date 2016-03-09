package main

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

func TestParseForCode(t *testing.T) {
	var tests = []struct {
		Answer    string
		InputHTML string
	}{
		{
			"Happy Code",
			"<body><div>Foo Bar <b>foobar</b> abc <code>Happy Code</code></div></body>",
		},
		{
			"Happy Pre",
			"<body><div>Foo Bar <b>foobar</b> abc <pre>Happy Pre</pre></div>",
		},
		{
			"Happy Code",
			"<body><div>Foo Bar <b>foobar</b> abc <pre><code>Happy Code</code></pre></div></body>",
		},
		{
			"",
			"<body><div>Foo Bar <b>foobar</b></div></body>",
		},
	}

	for _, s := range tests {
		aw := &answer{}
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(s.InputHTML))
		aw.html = doc.Find("body").First()

		aw.parseForCode()

		assert.Equal(t, s.Answer, aw.shortText)
	}
}

func TestParseForText(t *testing.T) {
	var tests = []struct {
		Answer    string
		InputHTML string
	}{
		{
			"Foo Bar foobar abc Happy Code",
			"<body><div class=\"post-text\">Foo Bar <b>foobar</b> abc <code>Happy Code</code></div></body>",
		},
		{
			"",
			"<body><div>Foo Bar <b>foobar</b> abc <code>Happy Code</code></div></body>",
		},
	}

	for _, s := range tests {
		aw := &answer{}
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(s.InputHTML))
		aw.html = doc.Find("body").First()

		aw.parseForText()

		assert.Equal(t, s.Answer, aw.longText)
	}
}

func TestParseForVotes(t *testing.T) {
	var tests = []struct {
		Votes     int
		InputHTML string
	}{
		{
			-300,
			"<body><div class=\"vote-count-post\">-300</div></body>",
		},
		{
			211,
			"<body><div class=\"vote-count-post\">211</div></body>",
		},
		{
			0,
			"<body><div>Foo Bar <b>foobar</b> abc <code>Happy Code</code></div></body>",
		},
	}

	for _, s := range tests {
		aw := &answer{}
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(s.InputHTML))
		aw.html = doc.Find("body").First()

		aw.parseForVotes()

		assert.Equal(t, s.Votes, aw.votes)
	}
}

func TestCreateGoogleSearchURL(t *testing.T) {
	var tests = []struct {
		Expected string
		Query    string
		Site     string
	}{
		{
			"https://www.google.com/search?q=site:stackoverflow.com+windows+get+date+command+line",
			"windows get date command line",
			"stackoverflow.com",
		},
		{
			"https://www.google.com/search?q=site:stackoverflow.com+" + url.QueryEscape("\"foobar\" for $$%%"),
			"\"foobar\" for $$%%",
			"stackoverflow.com",
		},
	}

	for _, s := range tests {
		assert.Equal(t, s.Expected, createGoogleSearchURL(s.Query, s.Site))
	}
}
