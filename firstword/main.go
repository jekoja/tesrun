package main

import "fmt"

func FirstWord(s string) string {
	word := ""
	for _, i := range s {
		if i == ' ' || i == '\t' || i == '\n' {
			break
		}
		word += string(i)
	}
	return word + "\n"
}
func main() {
	fmt.Print(FirstWord("Hello World"))
}
