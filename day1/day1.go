package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	leftList := make([]int, 0)
	rightList := make([]int, 0)

	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		lr := strings.Fields(line)
		if len(lr) != 2 {
			panic(len(lr))
		}
		l, err := strconv.Atoi(lr[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(lr[1])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, l)
		rightList = append(rightList, r)
	}

	occurances := make(map[int]int)

	for _, v := range leftList {
		if _, ok := occurances[v]; !ok {
			occurances[v] = 0
		}
	}

	for _, v := range rightList {
		if c, ok := occurances[v]; ok {
			occurances[v] = c + 1
		}
	}

	var score int

	for k, v := range occurances {
		score += k * v
	}

	fmt.Println(score)
}
