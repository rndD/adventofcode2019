package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	inputBytes, err := ioutil.ReadFile("./input.txt")
	check(err)
	input := string(inputBytes)
	numbers := strings.Split(input, "\n")

	// fmt.Print(input)

	sum := 0

	for _, i := range numbers {
		j, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
			continue
		}
		z := (int(j) / 3) - 2
		sum += z
	}
	fmt.Println(sum)
}
