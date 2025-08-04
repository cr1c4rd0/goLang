package main

import (
	"fmt"
	"strconv"
)

func main() {
	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(int32(integer16) + integer32)

	// Conversion from string to int
	s := "100"
	i, _ := strconv.Atoi(s)
	fmt.Println((i + i))

	// Conversion from int to string
	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s)
}
