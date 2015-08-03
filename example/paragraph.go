package main

import (
	"fmt"
	"net/http"

	"github.com/linkosmos/tokeq"
	"golang.org/x/net/html"
)

func main() {
	response, err := http.Get("https://golang.org/pkg/fmt/")
	if err != nil {
		fmt.Errorf("http.Get error: %s", err)
		return
	}
	toks := tokeq.Toks{
		tokeq.Tok{
			Match: tokeq.MatchP,
			Callback: func(n *html.Node) {
				fmt.Println(tokeq.FindText(n))
			},
		},
	}

	err = tokeq.ParseResponseWithDefer(toks, response)
	if err != nil {
		fmt.Errorf("tokeq.ParseResponseWithDefer error: %s", err)
	}
}
