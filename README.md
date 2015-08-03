# tokeq
Match &amp; Callback iterator for html.Node - aimed for performance &amp; HTML document reusability

[![Build Status](https://travis-ci.org/linkosmos/tokeq.svg?branch=master)](https://travis-ci.org/linkosmos/tokeq)
[![GoDoc](http://godoc.org/github.com/linkosmos/tokeq?status.svg)](http://godoc.org/github.com/linkosmos/tokeq)
[![BSD3License](http://img.shields.io/badge/license-BSD3-blue.svg)](http://opensource.org/licenses/BSD-3-Clause)

## Installation

```go
go get github.com/linkosmos/tokeq
```

## Godeps

```go
godep get github.com/linkosmos/tokeq
```

## Usage

```go
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
		fmt.Printf("tokeq.ParseResponseWithDefer error: %s", err)
	}
}
```

## Contributing

1. Fork it ( https://github.com/linkosmos/tokeq/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add tokeq feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
