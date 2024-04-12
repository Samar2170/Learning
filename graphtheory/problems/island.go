package main

import (
	"fmt"
	"strconv"
)

func islandCount(grid [][]string) int {
	visited := make(map[string]struct{})
	count := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if explore(grid, r, c, visited) {
				count++
			}
		}
	}
	return count
}

func explore(grid [][]string, row, col int, visited map[string]struct{}) bool {
	rowInbound := row >= 0 && row < len(grid)
	colInbound := col >= 0 && col < len(grid[0])
	if (!rowInbound) || (!colInbound) {
		return false
	}
	if grid[row][col] != "L" {
		return false
	}
	lookupStr := strconv.Itoa(row) + "," + strconv.Itoa(col)
	if _, ok := visited[lookupStr]; ok {
		return false
	}
	visited[lookupStr] = struct{}{}
	fmt.Println(visited)
	explore(grid, row-1, col, visited)
	explore(grid, row+1, col, visited)
	explore(grid, row, col+1, visited)
	explore(grid, row, col-1, visited)
	return true
}

func main() {
	grid := [][]string{
		{"W", "L", "W", "W", "W"},
		{"W", "L", "L", "W", "W"},
		{"W", "W", "W", "L", "W"},
		{"W", "W", "L", "L", "W"},
		{"L", "W", "W", "L", "L"},
		{"L", "L", "W", "W", "W"},
	}
	println(islandCount(grid))
}
