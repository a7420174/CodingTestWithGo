package main

import (
	"bufio"
	"strconv"
	"strings"
	"fmt"
)

const INF int = 1e9

var input string = 
`4 7
1 2 4
1 4 6
2 1 3
2 3 7
3 1 5
3 4 4
4 3 2`

func unpack(s []int, vars... *int) {
	for i, element := range s {
		*vars[i] = element
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	/*
	출처: 이것이 취업을 위한 코딩 테스트다 with 파이썬 (나동빈 저)

	플로이드 워셜 알고리즘 예제
	 */
	sc := bufio.NewScanner(strings.NewReader(input))
	sc.Scan()
	n_m := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(n_m[0])
	m, _ := strconv.Atoi(n_m[1])
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, n+1)
		for j := range graph[i] {
			if i==j {
				graph[i][j] = 0
			} else {graph[i][j] = INF}
		}

	}

	for i:=0;i<m;i++ {
		a_b_c := generateSeqFromStdin(sc, 3)
		var a, b, c int
		unpack(a_b_c, &a, &b, &c)
		graph[a][b] = c
	}

	for i:=1;i<=n;i++ {
		for j:=1;j<=n;j++ {
			for k:=1;k<=n;k++ {
				graph[i][j] = min(graph[i][j], graph[i][k] + graph[k][j])
			}
		}
	}

	for i:=1;i<=n;i++ {
		for j:=1;j<=n;j++ {
			fmt.Print(graph[i][j], " ")
		}
		fmt.Println()
	}
}
