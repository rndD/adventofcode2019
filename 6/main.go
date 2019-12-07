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

// type Node struct {
// 	name string
// 	children []*Node
// }
type nodesMap map[string][]string

func follow(nodes nodesMap, n int, current string) int {
	if current != "COM" {
		n++
	}
	if _, ok := nodes[current]; !ok {
		fmt.Println("in " + current + " " + strconv.Itoa(n))
		return n
	}
	z := 0
	for _, name := range nodes[current] {

		z += follow(nodes, n, name)

	}

	return z + n
}

func main() {

	inputBytes, err := ioutil.ReadFile("./input.txt")
	check(err)

	input := string(inputBytes)
	connections := strings.Split(input, "\n")

	nodes := make(nodesMap)
	// numbers := make([]int, len(strings))
	for _, connection := range connections {
		c := strings.Split(connection, ")")
		nodeName1 := c[0]
		nodeName2 := c[1]

		if _, ok := nodes[nodeName1]; !ok {
			nodes[nodeName1] = make([]string, 0)
		}
		nodes[nodeName1] = append(nodes[nodeName1], nodeName2)
	}

	c := follow(nodes, 0, "COM")
	fmt.Println(c)

}
