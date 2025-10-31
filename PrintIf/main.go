package main

import "fmt"

func PrintIf(s string) string {
	if s == "" {
		return "G\n"
	}
	if len(s) >= 3 {
		return "G\n"
	}
	return "invalid input\n"
}

func main() {
	fmt.Print(PrintIf(""))      // G
	fmt.Print(PrintIf("Hi"))    // invalid input
	fmt.Print(PrintIf("Hello")) // G
}
