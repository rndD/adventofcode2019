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
	finalImage := make([]int, maxY*maxX)

	for i := range layers[0] {
		// curPixelValue := 2
		for _, l := range layers {
			if l[i] == 2 {
				continue
			}
			finalImage[i] = l[i]
			break
		}
	}
	for i := 0; i < maxX*maxY; i += maxX {
		fmt.Println(finalImage[i : i+maxX])
	}

}
