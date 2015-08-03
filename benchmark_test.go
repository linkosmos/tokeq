package tokeq

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func BenchmarkDissectNodesFindP(t *testing.B) {
	input := strings.NewReader(benchTestHTML)
	toks := Toks{
		Tok{
			Match:    MatchP,
			Callback: func(n *html.Node) {},
		},
	}
	document, _ := html.Parse(input)
	for n := 0; n < t.N; n++ {
		DissectNodes(toks, document)
	}
}
