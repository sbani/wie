package main

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "strings"
    "github.com/PuerkitoBio/goquery"
)

func TestGetCodeFromAnswer(t *testing.T) {
    var tests = []struct{
    	Answer string
    	HasAnswer bool
    	InputHTML string
    }{
    	{
    	    "Happy Code",
    	    true,
    	    "<div>Foo Bar <b>foobar</b> abc <code>Happy Code</code></div>",
    	},
    	{
    	    "Happy Pre",
    	    true,
    	    "<div>Foo Bar <b>foobar</b> abc <pre>Happy Pre</pre></div>",
    	},
    	{
    	    "Happy Code",
    	    true,
    	    "<div>Foo Bar <b>foobar</b> abc <pre><code>Happy Code</code></pre></div>",
    	},
    	{
    	    "",
    	    false,
    	    "<div>Foo Bar <b>foobar</b></div>",
    	},
    }

    for _, s := range tests {
    	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(s.InputHTML))
        result, found := getCodeFromAnswer(doc.Find("div").First())

        assert.Equal(t, s.HasAnswer, found)
        assert.Equal(t, s.Answer, result)
    }
}

func TestGetCompleteAnswer(t *testing.T) {
    var tests = []struct{
    	Answer string
    	HasAnswer bool
    	InputHTML string
    }{
    	{
    	    "Foo Bar foobar abc Happy Code",
    	    true,
    	    "<body><div class=\"post-text\">Foo Bar <b>foobar</b> abc <code>Happy Code</code></div></body>",
    	},
    	{
    	    "",
    	    false,
    	    "<body><div>Foo Bar <b>foobar</b> abc <code>Happy Code</code></div></body>",
    	},
    }

    for _, s := range tests {
    	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(s.InputHTML))
        result, found := getCompleteAnswer(doc.Find("body").First())

        assert.Equal(t, s.HasAnswer, found)
        assert.Equal(t, s.Answer, result)
    }
}
