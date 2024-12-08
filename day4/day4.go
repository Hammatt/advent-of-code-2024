package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid := loadGrid(b)
	count := 0

	for i := 1; i < len(grid)-1; i++ {
		row := grid[i]
		for j := 1; j < len(row)-1; j++ {
			r := grid[i][j]
			if r == 'A' {
				var b strings.Builder
				b.WriteRune(grid[i-1][j-1])
				b.WriteRune(grid[i-1][j+1])
				b.WriteRune(grid[i+1][j+1])
				b.WriteRune(grid[i+1][j-1])

				if b.String() == "MMSS" || b.String() == "MSSM" || b.String() == "SSMM" || b.String() == "SMMS" {
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
