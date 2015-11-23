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
		fmt.Printf("http.Get error: %s", err)
		return
	}

	err = tokeq.ParseResponseWithDefer(response, tokeq.Tok{
		Match: tokeq.MatchP,
		Callback: func(n *html.Node) {
			fmt.Println(tokeq.FindText(n))
		},
	})
	if err != nil {
		fmt.Printf("tokeq.ParseResponseWithDefer error: %s", err)
	}
}
