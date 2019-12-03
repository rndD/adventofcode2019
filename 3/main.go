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
	lines := strings.Split(input, "\n")
	// len := 100
	// middle := len / 2
	// field := [100][100]uint8{}

	for _, line := range lines {
		commands := strings.Split(line, ",")
		for _, command := range commands {
			arg := command[:1]
			count, _ := strconv.Atoi(command[1:])
			fmt.Println(arg, count)
		}

	}
	// initNumbers := make([]int, len(strings))

	// fmt.Print(input)

	// for i, s := range strings {
	// 	initNumbers[i], _ = strconv.Atoi(s)
	// }

	fmt.Println("end")

}
