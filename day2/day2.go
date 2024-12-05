package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type report struct {
	levels []int
}

func main() {
	reports := make([]report, 0)

	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		rl := scanner.Text()

		l := strings.Fields(rl)
		levels := make([]int, 0)
		for _, v := range l {
			level, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			levels = append(levels, level)
		}

		reports = append(reports, report{levels})
	}

	c := 0
	for _, r := range reports {
		if r.isSafe() {
			c++
		}
	}

	fmt.Println(c)
}

func (r report) isSafe() bool {
	i := 0
	if r.levels[0] > r.levels[1] {
		slices.Reverse(r.levels)
	}
	p := r.levels[0]
	for i < len(r.levels)-1 {
		i++

		c := max(r.levels[i], -r.levels[i])

		if c < p {
			return false
		}

		diff := c - p
		if diff > 3 || diff < 1 {
			return false
		}

		p = c
	}

	return true
}
