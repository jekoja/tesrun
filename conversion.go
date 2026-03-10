package main

import (
	"fmt"
	"strconv"
)

func hexToDecimal(hexStr string) (int64, error) {

	decimal, err :=
		strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return 0, err
	}
	return decimal, nil
}

func main() {
	fmt.Println(hexToDecimal("1E"))
	fmt.Println(hexToDecimal("FF"))
}
