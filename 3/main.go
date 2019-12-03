package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func print2D(a *[30000][30000]int) {
	for _, line := range a {
		fmt.Println(line)
	}
}

func setLine(nLine int, a *[30000][30000]int, x int, y int) {
	if a[x][y] < nLine {
		a[x][y] += nLine
	}
}

func findSmallestDist(a *[30000][30000]int, centerX int, centerY int) int {
	smallestDistance := 0.0
	for y, lineX := range a {
		for x, value := range lineX {
			if value == 3 {
				dist := math.Sqrt(math.Pow(float64(x-centerX), 2)) + math.Sqrt(math.Pow(float64(y-centerY), 2))
				if smallestDistance == 0 || dist < smallestDistance {
					smallestDistance = dist
				}
			}
		}
	}
	return int(smallestDistance)
}

func main() {
	inputBytes, err := ioutil.ReadFile("./input.txt")
	check(err)
	input := string(inputBytes)
	lines := strings.Split(input, "\n")
	len := 30000
	middle := len / 2
	field := [30000][30000]int{}

	for nLine, line := range lines {
		nLine++
		commands := strings.Split(line, ",")
		x := middle
		y := middle
		for _, command := range commands {
			arg := command[:1]
			count, _ := strconv.Atoi(command[1:])
			if arg == "R" {
				for i := 0; i < count; i++ {
					x++
					setLine(nLine, &field, x, y)
				}
			}
			if arg == "L" {
				for i := 0; i < count; i++ {
					x--
					setLine(nLine, &field, x, y)
				}
			}
			if arg == "D" {
				for i := 0; i < count; i++ {
					y++
					setLine(nLine, &field, x, y)
				}
			}
			if arg == "U" {
				for i := 0; i < count; i++ {
					y--
					setLine(nLine, &field, x, y)
				}
			}

		}

	}
	// initNumbers := make([]int, len(strings))

	// fmt.Print(input)

	// for i, s := range strings {
	// 	initNumbers[i], _ = strconv.Atoi(s)
	// }

	// print2D(&field)
	fmt.Println(findSmallestDist(&field, middle, middle))

}
