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

func calc(j int, result int) int {
	z := (j / 3) - 2
	result += z

	if z > 6 {
		return calc(z, result)
	}

	return result

}

func main() {
	inputBytes, err := ioutil.ReadFile("./input.txt")
	check(err)
	input := string(inputBytes)
	strings := strings.Split(input, ",")
	initNumbers := make([]int, len(strings))
	i := 0

	// fmt.Print(input)

	for i, s := range strings {
		initNumbers[i], _ = strconv.Atoi(s)
	}
	numbers := make([]int, len(initNumbers))

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			copy(numbers, initNumbers)
			numbers[1] = x
			numbers[2] = y
			i = 0

			for true {
				if numbers[i] == 1 {
					to := numbers[i+3]
					from1 := numbers[i+1]
					from2 := numbers[i+2]
					numbers[to] = numbers[from1] + numbers[from2]
					i += 4
					continue
				}

				if numbers[i] == 2 {

					to := numbers[i+3]
					from1 := numbers[i+1]
					from2 := numbers[i+2]
					numbers[to] = numbers[from1] * numbers[from2]
					i += 4
					continue
				}

				if numbers[i] == 99 {
					if numbers[0] == 19690720 {
						fmt.Println(numbers)
						fmt.Println(100*numbers[1] + numbers[2])
					} else {
						// fmt.Println(numbers[0])
					}
					break
				}

				panic("Strange op code: " + string(numbers[i]))
			}
		}
	}

	fmt.Println("end")

}
