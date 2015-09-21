package tokeq

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Matcher - used in Tok
type Matcher func(tt html.NodeType, a atom.Atom) bool

// MatcherCallback - uses in Tok as a callback when match occurs.
// Contents of the (t *html.Token) may change on the next call to Next.
type MatcherCallback func(t *html.Node)

// Tok - contains html.Node Matcher & Callback
type Tok struct {
	Match    Matcher
	Callback MatcherCallback
}

// Toks - array of Tok's
type Toks []Tok

// Add - add new Tok to Toks
func (toks *Toks) Add(tok ...Tok) {
	(*toks) = append((*toks), tok...)
}

// Iterate - iterate through tok's and if there is a match send *token to calback
func (toks *Toks) Iterate(t *html.Node) {
	for _, tok := range *toks {
		if tok.Match(t.Type, t.DataAtom) {
			tok.Callback(t)
		}
	}
}
