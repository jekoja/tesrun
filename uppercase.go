package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "chrIs"
	result := strings.ToUpper(name[:1]) + strings.ToLower(name[1:])

	fmt.Println(result)
}
