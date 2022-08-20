package main

import (
	"bufio"
	"strconv"
	"strings"
)

func generateGridFromStdin(inputScanner *bufio.Scanner, n, m int) [][]int {
	grid := make([][]int, m)
	for i := range grid {
	    grid[i] = make([]int, n)
	}
	
	i := 0

	for inputScanner.Scan() {

		j := 0

		lineScanner := bufio.NewScanner(strings.NewReader(inputScanner.Text()))
		lineScanner.Split(bufio.ScanWords)

		for lineScanner.Scan() {

			grid[i][j], _ = strconv.Atoi(lineScanner.Text())
			j++

		}

		i++

		if i == m {
			break
		}
	}

	return grid
}
