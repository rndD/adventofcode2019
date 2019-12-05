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
func isZero(s byte) bool {
	return s == 48
}

func parse(n int) (opcode int, posmod1 bool, posmod2 bool, posmod3 bool) {
	s := strconv.Itoa(n)
	// fmt.Println(s)

	if (len(s) == 1) || (len(s) == 2) {
		return n, true, true, true
	}

	// 11101
	// 01234
	if len(s) == 3 {
		i, _ := strconv.Atoi(string(s[1:]))
		return i, isZero(s[0]), true, true
	}

	if len(s) == 4 {
		i, _ := strconv.Atoi(string(s[2:]))
		return i, isZero(s[1]), isZero(s[0]), true
	}

	if len(s) == 5 {
		i, _ := strconv.Atoi(string(s[3:]))
		return i, isZero(s[2]), isZero(s[1]), isZero(s[0])
	}

	panic("whaa")
}

func main() {

	inputBytes, err := ioutil.ReadFile("./input.txt")
	check(err)

	input := string(inputBytes)
	strings := strings.Split(input, ",")
	numbers := make([]int, len(strings))

	// fmt.Print(input)

	for i, s := range strings {
		numbers[i], _ = strconv.Atoi(s)
	}
	i := 0

	for true {
		opcode, posmod1, posmod2, _ := parse(numbers[i])
		// fmt.Println("op", opcode, numbers[i])
		if opcode == 1 {
			to := numbers[i+3]
			var a, b int

			if posmod1 {
				from := numbers[i+1]
				a = numbers[from]
			} else {
				a = numbers[i+1]
			}

			if posmod2 {
				from := numbers[i+2]
				b = numbers[from]
			} else {
				b = numbers[i+2]
			}

			numbers[to] = a + b
			i += 4
			continue
		}

		if opcode == 2 {
			to := numbers[i+3]
			var a, b int

			if posmod1 {
				from := numbers[i+1]
				a = numbers[from]
			} else {
				a = numbers[i+1]
			}

			if posmod2 {
				from := numbers[i+2]
				b = numbers[from]
			} else {
				b = numbers[i+2]
			}

			numbers[to] = a * b
			i += 4
			continue
		}

		if opcode == 3 {
			input := 1
			to := numbers[i+1]
			numbers[to] = input
			i += 2
			continue
		}

		if opcode == 4 {
			from := numbers[i+1]
			fmt.Println("Output", numbers[from])
			i += 2
			continue
		}

		if opcode == 99 {
			// if numbers[0] == 19690720 {
			// 	fmt.Println(numbers)
			// 	fmt.Println(100*numbers[1] + numbers[2])
			// } else {
			// 	// fmt.Println(numbers[0])
			// }
			if numbers[i-2] == 4 {
				fmt.Println("good end")
				break
			}

			fmt.Println("bad end")
			break
		}

		panic("Strange op code: " + string(numbers[i]) + " i:" + string(i))
	}

}
