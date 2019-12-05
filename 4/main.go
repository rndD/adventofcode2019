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

func itHasAdjacentNumberButNotMajor(s string) bool {
	doubles := make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			n := string(s[i])

			if _, ok := doubles[n]; !ok {
				doubles[n] = 0
			}

			doubles[n]++
		}
	}

	doublesN := make([]string, len(doubles))

	i := 0
	for k := range doubles {
		doublesN[i] = k
		i++
	}

	if len(doublesN) == 3 {
		return true
	}

	if len(doublesN) == 2 {
		if (doubles[doublesN[0]] == 2) && (doubles[doublesN[1]] == 2) {
			return false
		}
		return true
	}

	if len(doublesN) == 1 {
		if doubles[doublesN[0]] == 1 {
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

		if !itHasAdjacentNumberButNotMajor(s) {
			continue
		}

		fmt.Println(i)
		result++
	}

	fmt.Println(result)
	fmt.Println(itHasAdjacentNumberButNotMajor("111122"))
}
