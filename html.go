package tokeq

import (
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
)

// IsTextHTML - compare header content type
func IsTextHTML(contentType string) bool {
	// mediatype, _, err := mime.ParseMediaType(contentType)
	// if err != nil {
	// 	return false
	// }
	// return mediatype == "text/html"
	return strings.Contains(contentType, "text/html")
}

// FindNodes - recursively find nodes
func FindNodes(input *html.Node, match Matcher, callback MatcherCallback) {
	for c := input.FirstChild; c != nil; c = c.NextSibling {
		if match(c.Type, c.DataAtom) {
			callback(c)
		}
		FindNodes(c, match, callback)
	}
}

// ParseResponse - wrapps sequence of URL fate functions
// user is response to handle: defer response.Body.Close()
func ParseResponse(response *http.Response, toks ...Tok) error {
	contentType := response.Header.Get("Content-Type")
	if !IsTextHTML(contentType) {
		return ErrResponseBodyIsNotHTML
	}
	if response.Body == nil {
		return ErrResponseBodyIsEmpty
	}
	r, err := charset.NewReader(response.Body, contentType)
	if err != nil {
		return err
	}
	return ParseReader(r, toks...)
}

// ParseResponseWithDefer - same as ParseResponse but with defer response.Body.Close()
func ParseResponseWithDefer(response *http.Response, toks ...Tok) error {
	contentType := response.Header.Get("Content-Type")
	if !IsTextHTML(contentType) {
		return ErrResponseBodyIsNotHTML
	}
	if response.Body == nil {
		return ErrResponseBodyIsEmpty
	}
	defer response.Body.Close()
	r, err := charset.NewReader(response.Body, contentType)
	if err != nil {
		return err
	}
	return ParseReader(r, toks...)
}

// ParseReader - parses io.Reader, expected input is HTML page
func ParseReader(input io.Reader, toks ...Tok) error {
	document, err := html.Parse(input)
	if err != nil {
		return err
	}
	DissectNodes(document, toks...)
	return nil
}

// DissectNodes - range toks through recursively through FindNodes
func DissectNodes(input *html.Node, toks ...Tok) {
	for _, tok := range toks {
		FindNodes(input, tok.Match, tok.Callback)
	}
}
