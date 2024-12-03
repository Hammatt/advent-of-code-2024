package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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

	slices.Sort(leftList)
	slices.Sort(rightList)

	td := 0
	for i := range len(leftList) {
		diff := leftList[i] - rightList[i]
		td += int(math.Abs(float64(diff)))
	}

	fmt.Println(td)
}
