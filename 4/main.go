package main

import (
	"fmt"
	"strconv"
)

func itHasAdjacentNumber(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func itDecrising(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if int(s[i]) > int(s[i+1]) {
			return true
		}
	}
	return false
}

func main() {
	from := 134792
	to := 675810
	result := 0

	for i := from; i <= to; i++ {
		s := strconv.Itoa(i)
		if !itHasAdjacentNumber(s) {
			continue
		}

		if itDecrising(s) {
			// fmt.Println("descrising", s)
			continue
		}

		result++
	}

	fmt.Println(result)
}
