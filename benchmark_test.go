package tokeq

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func BenchmarkDissectNodesFindP(t *testing.B) {
	input := strings.NewReader(benchTestHTML)
	tok := Tok{
		Match:    MatchP,
		Callback: func(n *html.Node) {},
	}
	document, _ := html.Parse(input)
	for n := 0; n < t.N; n++ {
		DissectNodes(document, tok)
	}
}
