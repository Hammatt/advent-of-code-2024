package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)

	regex := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")
	matches := regex.FindAllString(input, -1)

	answer := 0

	for _, match := range matches {
		argsStr := match[4 : len(match)-1]
		args := strings.Split(argsStr, ",")
		l, _ := strconv.Atoi(args[0])
		r, _ := strconv.Atoi(args[1])

		answer += l * r
	}

	fmt.Println(answer)
}
