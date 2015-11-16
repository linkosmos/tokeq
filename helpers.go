package tokeq

import (
	"fmt"

	"golang.org/x/net/html"
)

// PrintNode - convenient to use as Tok.Callback to see html.Node dump
func PrintNode(n *html.Node) {
	fmt.Println(PrettyPrint(n))
}

// PrintNodes - convenient to use as Tok.Callback to see html.Node siblings & parents
func PrintNodes(n *html.Node) {
	fmt.Println(PrettyPrint(n))
	fmt.Println("Parent", PrettyPrint(n.Parent))
	fmt.Println("FirstChild", PrettyPrint(n.FirstChild))
	fmt.Println("LastChild", PrettyPrint(n.LastChild))
	fmt.Println("PrevSibling", PrettyPrint(n.PrevSibling))
	fmt.Println("NextSibling", PrettyPrint(n.NextSibling))
}

// PrettyPrint - pretty print node
func PrettyPrint(n *html.Node) string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf(
		`
Type: %s
DataAtom: %s
Data: %s
Namespace: %s
Attr: %s
`, n.Type, n.DataAtom, n.Data, n.Namespace, n.Attr)
}
