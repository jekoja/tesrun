package main

import (
	"fmt"
	"strings"
)

func replace(s string) string {
	data := strings.ReplaceAll(s, " ,", ",")
	return data
}

func main() {
	val := "hello , how are you"
	fmt.Println(replace(val))
}
