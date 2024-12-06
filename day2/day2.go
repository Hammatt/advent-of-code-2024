package main

import (
	"bufio"
	"fmt"
	"os"
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
		} else {
			orig := make([]int, len(r.levels))
			copy(orig, r.levels)
			for i := range orig {
				r.levels = make([]int, len(orig))
				copy(r.levels, orig)
				r.levels = append(r.levels[:i], r.levels[i+1:]...)
				if r.isSafe() {
					c++
					break
				}
			}
		}
	}

	fmt.Println(c)
}

func (r report) isSafe() bool {
	i := 0

	comparitor := lessThan
	if r.levels[0] < r.levels[1] {
		comparitor = greaterThan
	}
	p := r.levels[0]
	for i < len(r.levels)-1 {
		i++

		c := r.levels[i]

		if comparitor(p, c) {
			return false
		}

		diff := max(c, p) - min(c, p)
		if diff > 3 || diff < 1 {
			return false
		}

		p = c
	}

	return true
}

func lessThan(a, b int) bool {
	return a < b
}

func greaterThan(a, b int) bool {
	return a > b
}
