package tokeq

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func BenchmarkGoQueryFindP(t *testing.B) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(benchTestHTML))

	for n := 0; n < t.N; n++ {
		doc.Find("p").Each(func(i int, s *goquery.Selection) {
		})
	}
}
