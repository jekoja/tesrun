package main

import "fmt"

func Rectperimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}
	return 2 * (w + h)
}
func main() {
	fmt.Println(Rectperimeter(1, -2))
}
