package main

import "fmt"

func RetainFirstHalf(str string) string {
	length := len(str)
	if length == 0 {
		return ""
	}

	if length == 1 {
		return str
	}

	half := length / 2
	return str[:half]
}
func main() {
	fmt.Println(RetainFirstHalf("cluster"))
	fmt.Println(RetainFirstHalf("leefcenter"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("missionary"))

}
