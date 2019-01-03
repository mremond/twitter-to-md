package metadata

import (
	"io"

	"golang.org/x/net/html"
)

// Dublin Core metadata are defined here: http://www.dublincore.org/documents/dces/
// Open Graph metadata (Facebook) are defined here: http://ogp.me/
// Twitter metadata are defined here: https://developer.twitter.com/en/docs/tweets/optimize-with-cards/overview/markup

type Properties map[string]string

type Page struct {
	Properties Properties
}

// GetTitle returns the page title based on defined priorities (dc > og > twitter > title)
func (p Page) GetTitle() string {
	propNames := []string{"dc.title", "og:title", "twitter:title", "title"}
	for _, name := range propNames {
		value := p.Properties[name]
		if value != "" {
			return value
		}
	}
	return ""
}

var knownProperties = []string{
	// Dublin Core (older version)
	"dc.title",
	// Dublin Core (HTML 5)
	"dc:title", "dc:creator",
	// Open Graph
	"og:title", "og:type", "og:url", "og:image",
	"og:description", "og:site_name",
	// Twitter
	"twitter:card", "twitter:site", "twitter:title",
	"twitter:image",
}

func FromReader(body io.Reader) (Page, error) {
	var p Page
	p.Properties = make(map[string]string)

	tokenizer := html.NewTokenizer(body)
Loop:
	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			err := tokenizer.Err()
			if err != io.EOF {
				return p, err
			}
			break Loop

		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			switch token.Data {
			case "meta":
				meta := extractMeta(token)
				if contains(knownProperties, meta.property) {
					p.Properties[meta.property] = meta.content
				}
			case "title":
				//the next token should be the page title
				tokenType = tokenizer.Next()
				if tokenType == html.TextToken {
					// Use page title but keep on searching of open graph title, which is often more accurate
					p.Properties["title"] = tokenizer.Token().Data
				}
			}
		case html.EndTagToken:
			token := tokenizer.Token()
			if token.Data == "head" {
				// We finished processing HTML head, no more metadata expected.
				break Loop
			}
		}
	}

	return p, nil
}

type meta struct {
	property string
	content  string
}

func extractMeta(token html.Token) meta {
	var m meta

	for _, attr := range token.Attr {
		if attr.Key == "property" {
			m.property = attr.Val
		}
		// Twitter is incorrectly using name attribute to hold metadata
		// For details, see: https://www.ctrl.blog/entry/rdfa-socialmedia-metadata
		if m.property == "" && attr.Key == "name" {
			m.property = attr.Val
		}
		if attr.Key == "content" {
			m.content = attr.Val
		}
	}
	return m
}

// TODO also extract og:image. e.g.:
// <meta property="og:image" content="https://gigaom.com/wp-content/uploads/sites/1/2011/01/sonosgroup-804x516.jpg" />
// TODO extract dcterms.title
//  example: <meta name='dcterms.title' content='Amazon&#8217;s dead serious about the enterprise cloud' />
//  on: https://gigaom.com/2012/11/21/amazons-dead-serious-about-the-enterprise-cloud/

//============================================================================
// Helper functions

func contains(array []string, str string) bool {
	for _, elt := range array {
		if elt == str {
			return true
		}
	}
	return false
}
