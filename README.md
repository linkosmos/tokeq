# tokeq
Match &amp; Callback iterator for html.Node - aimed for performance &amp; HTML document reusability

[![Build Status](https://travis-ci.org/linkosmos/tokeq.svg?branch=master)](https://travis-ci.org/linkosmos/tokeq)
[![GoDoc](http://godoc.org/github.com/linkosmos/tokeq?status.svg)](http://godoc.org/github.com/linkosmos/tokeq)
[![BSD3License](http://img.shields.io/badge/license-BSD3-blue.svg)](http://opensource.org/licenses/BSD-3-Clause)

## Benchmark

```go
go test -run=10000 -bench=. -benchmem
```

Finding html <p> paragraph                 | operations | time / operation | Bytes / operation | allocations / operation
------------------------------------------ | ---------- | ---------------- | ----------------- | ------------------------
BenchmarkGoQueryFindP                      |     5000   |   277185 ns/op   |  46059 B/op       |   929 allocs/op
BenchmarkStandardLibraryTokenFindP         |      500   |  3698731 ns/op   | 291695 B/op       |  9240 allocs/op
BenchmarkStandardLibraryNodeFindP          |    20000   |    78340 ns/op   |     64 B/op       |     0 allocs/op
BenchmarkDissectNodesFindP                 |    20000   |    68640 ns/op   |     64 B/op       |     0 allocs/op

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
```

## Contributing

1. Fork it ( https://github.com/linkosmos/tokeq/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add tokeq feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
