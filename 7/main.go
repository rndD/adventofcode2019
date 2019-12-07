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

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
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

	panic("whaat?")
}

func main() {

	inputBytes, err := ioutil.ReadFile("./input.txt")
	check(err)

	input := string(inputBytes)
	strings := strings.Split(input, ",")
	initNumbers := make([]int, len(strings))
	numbers := make([]int, len(strings))
	// fmt.Print(input)

	for i, s := range strings {
		initNumbers[i], _ = strconv.Atoi(s)
	}

	initSeq := []int{0, 1, 2, 3, 4}
	bigest := 0

	for _, seq := range permutation(initSeq) {
		fmt.Println("seq:", seq)
		userInput := 0
		for len(seq) > 0 {
			seqInput := 0
			seqInput, seq = seq[0], seq[1:]
			copy(numbers, initNumbers)
			i := 0
			firstInput := true

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
					to := numbers[i+1]
					if firstInput {
						numbers[to] = seqInput
						firstInput = false
					} else {
						numbers[to] = userInput
					}

					i += 2
					continue
				}

				if opcode == 4 {
					from := numbers[i+1]
					userInput = numbers[from]
					i += 2
					continue
				}

				if opcode == 5 {
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

					if a != 0 {
						i = b
					} else {
						i += 3
					}
					continue
				}

				if opcode == 6 {
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

					if a == 0 {
						i = b
					} else {
						i += 3
					}
					continue
				}

				if opcode == 7 {
					var a, b int
					to := numbers[i+3]

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

					if a < b {
						numbers[to] = 1
					} else {
						numbers[to] = 0
					}

					i += 4
					continue
				}

				if opcode == 8 {
					var a, b int
					to := numbers[i+3]

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

					if a == b {
						numbers[to] = 1
					} else {
						numbers[to] = 0
					}

					i += 4
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
						// fmt.Println("good end")
						break
					}

					fmt.Println("bad end")
					break
				}
				panic("Strange op code: " + strconv.Itoa(opcode) + ". i: " + strconv.Itoa(i))
			}
			if bigest < userInput {
				bigest = userInput
			}
		}
	}

	fmt.Println("Bigest: ", bigest)
}
