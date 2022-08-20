package main

import (
	"bufio"
	"strconv"
	"strings"
)

func generateSeqFromStdin(inputScanner *bufio.Scanner, n int) []int {
	seq := make([]int, n)

	i := 0

	for inputScanner.Scan() {
		
		lineScanner := bufio.NewScanner(strings.NewReader(inputScanner.Text()))
		lineScanner.Split(bufio.ScanWords)

		seq[i], _ = strconv.Atoi(lineScanner.Text())

		i++

		if i == n {
			break
		}
	}

	return seq
}


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
