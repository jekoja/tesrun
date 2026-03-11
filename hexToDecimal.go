package main

import (
	"fmt"
	"strconv"
)

func binToDecimal(binStr string) (int64, error) {

	decimal, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return 0, err
	}
	return decimal, nil
}

func main() {
	fmt.Println(binToDecimal("10"))
	fmt.Println(binToDecimal("1010"))
	fmt.Println(binToDecimal("11111111"))
}
