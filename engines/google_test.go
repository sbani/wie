package engines

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "net/url"
)

func TestCreateGoogle(t *testing.T) {
    var tests = []struct{
    	Expected Google
    	InputSSL bool
    }{
    	{
    		Google{
	    		"https://www.google.com",
	    		"search?q=site:%s+%s",
    		},
    		true,
    	},
    	{
    		Google{
	    		"http://www.google.com",
	    		"search?q=site:%s+%s",
    		},
    		false,
    	},
    }

    for _, s := range tests {
    	g:= s.Expected
    	created := CreateGoogle(s.InputSSL)

        assert.Equal(t, g.Host, created.Host)
        assert.Equal(t, g.URI, created.URI)
    }
}

func TestCreateSearchURL(t *testing.T) {
    var tests = []struct{
    	Expected string
    	Query string
    	Site string
    }{
    	{
    		"www.google.com/search?q=site:stackoverflow.com+windows+get+date+command+line",
    		"windows get date command line",
    		"stackoverflow.com",
    	},
    	{
    		"www.google.com/search?q=site:stackoverflow.com+" + url.QueryEscape("\"foobar\" for $$%%"),
    		"\"foobar\" for $$%%",
    		"stackoverflow.com",
    	},
    }

	// Without SSL
	g := CreateGoogle(false)
    for _, s := range tests {
        assert.Equal(t, "http://" + s.Expected, g.createSearchURL(s.Query, s.Site))
    }

    // With SSL
	gSSL := CreateGoogle(true)
    for _, s := range tests {
        assert.Equal(t, "https://" + s.Expected, gSSL.createSearchURL(s.Query, s.Site))
    }
}
