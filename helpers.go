package tokeq

import (
	"fmt"

	"golang.org/x/net/html"
)

// PrintNode - convenient to use as Tok.Callback to see html.Node dump
func PrintNode(n *html.Node) {
	fmt.Println(n)
	fmt.Println("Parent", n.Parent)
	fmt.Println("FirstChild", n.FirstChild)
	fmt.Println("LastChild", n.LastChild)
	fmt.Println("PrevSibling", n.PrevSibling)
	fmt.Println("NextSibling", n.NextSibling)
}
