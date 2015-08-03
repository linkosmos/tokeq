package tokeq

import (
	"io"
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func benchStandardLibraryToken(r io.Reader, match func(tt html.TokenType, a atom.Atom) bool, output func(html.Token)) {
	parser := html.NewTokenizer(r)
	for {
		tt := parser.Next() // TokenType
		if tt == html.ErrorToken {
			break
		}
		token := parser.Token()

		switch tt {
		case html.StartTagToken:
			if match(tt, token.DataAtom) {
				output(token)
			}
		case html.TextToken:
			if match(tt, token.DataAtom) {
				output(token)
			}
		case html.EndTagToken:
			if match(tt, token.DataAtom) {
				output(token)
			}
		case html.SelfClosingTagToken:
			if match(tt, token.DataAtom) {
				output(token)
			}
		case html.CommentToken:
			if match(tt, token.DataAtom) {
				output(token)
			}
		case html.DoctypeToken:
			if match(tt, token.DataAtom) {
				output(token)
			}
		}

	}
}

func BenchmarkStandardLibraryTokenFindP(t *testing.B) {
	for n := 0; n < t.N; n++ {
		benchStandardLibraryToken(strings.NewReader(benchTestHTML), func(t html.TokenType, a atom.Atom) bool {
			return t == html.StartTagToken && a == atom.P
		}, func(t html.Token) {})
	}
}

func benchStandardLibraryNode(input *html.Node, match func(n *html.Node) bool, output func(n *html.Node)) {
	for c := input.FirstChild; c != nil; c = c.NextSibling {
		if match(c) {
			output(c)
		}
		benchStandardLibraryNode(c, match, output)
	}
}

func BenchmarkStandardLibraryNodeFindP(t *testing.B) {
	document, _ := html.Parse(strings.NewReader(benchTestHTML))

	for n := 0; n < t.N; n++ {
		benchStandardLibraryNode(document, func(n *html.Node) bool {
			return n.Type == html.ElementNode && n.DataAtom == atom.P
		}, func(n *html.Node) {})
	}
}
