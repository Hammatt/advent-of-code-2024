package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(bytes.NewReader(b))
	updates, comp := parseInput(s)

	valid := 0
	sorted := 0
	for _, update := range updates {
		if isUpdateValid(update, comp) {
			v, _ := strconv.Atoi(update[len(update)/2])
			valid += v
		} else {
			slices.SortFunc(update, comp)
			v, _ := strconv.Atoi(update[len(update)/2])
			sorted += v
		}
	}

	fmt.Println(valid)
	fmt.Println(sorted)
}

func isUpdateValid(u []string, comp func(a, b string) int) bool {
	return slices.IsSortedFunc(u, comp)
}

func parseInput(s *bufio.Scanner) (updates [][]string, comp func(a, b string) int) {
	rules := make(map[string][]string)

	comp = func(a, b string) int {
		if slices.Contains(rules[a], b) {
			return -1
		}

		if slices.Contains(rules[b], b) {
			return 1
		}

		return 0
	}

	for s.Scan() {
		line := s.Text()
		if line == "" {
			updates = parseUpdates(s)
			return
		}

		rule := strings.Split(line, "|")
		if _, ok := rules[rule[0]]; ok {
			rules[rule[0]] = append(rules[rule[0]], rule[1])
		} else {
			constraint := make([]string, 0)
			constraint = append(constraint, rule[1])
			rules[rule[0]] = constraint
		}
	}

	return
}

func parseUpdates(s *bufio.Scanner) [][]string {
	updates := make([][]string, 0)
	for s.Scan() {
		line := s.Text()
		updates = append(updates, strings.Split(line, ","))
	}

	return updates
}
