package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type direction struct {
	y int
	x int
}

var tileSearchDirections = [...]direction{{-1, 0}, {0, -1}, {1, 0}, {0, 1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	answerRunes := []rune("XMAS")

	grid := loadGrid(b)
	count := 0

	for i := 0; i < len(grid); i++ {
		row := grid[i]
		for j := 0; j < len(row); j++ {
			for _, tsd := range tileSearchDirections {
				if tryTile(grid, i, j, answerRunes, 0, tsd) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func loadGrid(b []byte) [][]rune {
	grid := make([][]rune, 0)

	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		row := scanner.Text()
		runes := []rune(row)
		grid = append(grid, runes)
	}

	return grid
}

func tryTile(grid [][]rune, i, j int, word []rune, wordIndex int, tsd direction) bool {
	t := grid[i][j]
	w := word[wordIndex]

	if t != w {
		return false
	} else if t == w && wordIndex == len(word)-1 {
		return true
	}

	ni := i + tsd.y
	nj := j + tsd.x

	if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[ni]) && tryTile(grid, ni, nj, word, wordIndex+1, tsd) {
		return true
	}

	return false
}
