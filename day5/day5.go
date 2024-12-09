package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(bytes.NewReader(b))
	rules, updates := parseInput(s)

	valid := make([]int, 0)
	for _, update := range updates {
		if isUpdateValid(rules, update) {
			v, _ := strconv.Atoi(update[len(update)/2])
			valid = append(valid, v)
		}
	}

	ans := 0
	for _, v := range valid {
		ans += v
	}
	fmt.Println(ans)
}

func isUpdateValid(rules map[string][]string, u []string) bool {
	encountered := make(map[string]bool)
	for _, p := range u {
		if r, ok := rules[p]; ok {
			for _, constraint := range r {
				if _, enc := encountered[constraint]; enc {
					return false
				}
			}
		}
		encountered[p] = true
	}

	return true
}

func parseInput(s *bufio.Scanner) (rules map[string][]string, updates [][]string) {
	rules = make(map[string][]string)

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
