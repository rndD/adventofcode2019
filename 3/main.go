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

func print2D(a *[26000][26000][]int) {
	for _, line := range a {
		fmt.Println(line)
	}
}

func setLine(nLine int, a *[26000][26000][]int, x int, y int, aCache map[string]bool, distance int) {
	cacheKey := string(nLine) + ":" + string(x) + "," + string(y)
	if !aCache[cacheKey] {

		if a[x][y] == nil {
			z := make([]int, 0)
			a[x][y] = z
		}

		a[x][y] = append(a[x][y], distance)
		aCache[cacheKey] = true
	}

}

func findSmallestDist(a *[26000][26000][]int, centerX int, centerY int) int {
	smallestDistance := -1
	for _, lineX := range a {
		for _, value := range lineX {
			// fmt.Println(value)
			if len(value) == 2 {
				// fmt.Println(value)
				dist := value[0] + value[1]
				if smallestDistance == -1 || dist < smallestDistance {
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
	len := 26000
	middle := len / 2
	field := [26000][26000][]int{}
	filedCache := make(map[string]bool)

	for nLine, line := range lines {
		nLine++
		commands := strings.Split(line, ",")
		x := middle
		y := middle
		distance := 0
		for _, command := range commands {
			arg := command[:1]
			count, _ := strconv.Atoi(command[1:])
			if arg == "R" {
				for i := 0; i < count; i++ {
					x++
					distance++
					setLine(nLine, &field, x, y, filedCache, distance)
				}
			}
			if arg == "L" {
				for i := 0; i < count; i++ {
					x--
					distance++
					setLine(nLine, &field, x, y, filedCache, distance)
				}
			}
			if arg == "D" {
				for i := 0; i < count; i++ {
					y++
					distance++
					setLine(nLine, &field, x, y, filedCache, distance)
				}
			}
			if arg == "U" {
				for i := 0; i < count; i++ {
					y--
					distance++
					setLine(nLine, &field, x, y, filedCache, distance)
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
