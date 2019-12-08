package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findNum(a *[]int, findVal int) (count int) {
	for _, v := range *a {
		if v == findVal {
			count++
		}
	}
	return count
}

func main() {

	inputBytes, err := ioutil.ReadFile("./input.txt")
	check(err)

	input := string(inputBytes)
	layers := make([][]int, 0)
	layers = append(layers, make([]int, 0))
	curLayer := 0
	maxX := 25
	maxY := 6
	x := 0
	// y := 0
	for _, s := range input {
		if x == (maxX * maxY) {
			curLayer++
			layers = append(layers, make([]int, 0))
			x = 0
		}

		n, err := strconv.Atoi(string(s))
		check(err)
		layers[curLayer] = append(layers[curLayer], n)
		x++
	}

	fewestI := 0
	fewestCount := 1000
	for i, l := range layers {
		n := findNum(&l, 0)
		fmt.Println("len", len(l), n)
		if n < fewestCount {
			fewestI = i
			fewestCount = n
		}
	}

	fmt.Println(fewestI, len(layers))
	fmt.Println(findNum(&layers[fewestI], 1) * findNum(&layers[fewestI], 2))

}
