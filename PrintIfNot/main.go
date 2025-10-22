package main

import "fmt"

func PrintIfNot(s string) string {
	if len(s) >= 3 {
		return "invalid string\n"
	}
	return "G\n"
}
func main() {
	fmt.Println(PrintIfNot("aasdffg"))
	fmt.Println(PrintIfNot("ab"))
	fmt.Println(PrintIfNot("a"))
}
