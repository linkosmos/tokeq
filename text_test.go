package tokeq

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func ToNode(input string) *html.Node {
	n, err := html.ParseFragment(strings.NewReader(input), &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})
	if err != nil {
		panic(err)
	}
	return n[0]
}

var findTextTests = []struct {
	html     string
	expected string
}{
	{`<h1 class="details-list__item-title" itemprop="name">Emm <em>Again</em> Blue text span <span>Arti</span></h1>`, "Emm  Blue text span "},
	{`<h1 class="details-list__item-title" itemprop="name">Blue text span<span>Arti</span></h1>`, "Blue text span"},
	{`<h1 class="details-list__item-title" itemprop="name"><em>Again</em> Blue text span <span>Arti</span></h1>`, " Blue text span "},
	{`<h1>Blue text span</h1>`, "Blue text span"},
}

func TestFindText(t *testing.T) {
	for _, test := range findTextTests {
		n := ToNode(test.html)
		got := FindText(n)
		if got != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, got)
		}
	}
}

var findDeepTextTests = []struct {
	html     string
	expected string
}{
	{`<h1 class="details-list__item-title" itemprop="name">Emm <em>Again</em> Blue text span <span>Arti</span></h1>`, "Emm Again Blue text span Arti"},
	{`<h1 class="details-list__item-title" itemprop="name">Blue text span<span> Arti</span></h1>`, "Blue text span Arti"},
	{`<h1 class="details-list__item-title" itemprop="name"><em>Again</em> Blue text span <span>Arti</span></h1>`, "Again Blue text span Arti"},
	{`<h1>Blue text span</h1>`, "Blue text span"},
}

func TestDeepFindText(t *testing.T) {
	for _, test := range findDeepTextTests {
		n := ToNode(test.html)
		got := FindDeepText(n)
		if got != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, got)
		}
	}
}
