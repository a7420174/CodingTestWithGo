package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const INF int = 1e9

var input string = 
`7 9
1 2 29
1 5 75
2 3 35
2 6 34
3 4 7
4 6 23
4 7 13
5 6 53
6 7 25
`

type Path struct{
	x, y int
}

type BuildCost struct{
	c int
	p Path
}

func main() {
	/*
	출처: 이것이 코딩 테스트다 with 파이썬 - 나동빈 저

	크루스칼 알고리즘

	도로 최소 건설비용 구하기
	 */
	sc := bufio.NewScanner(strings.NewReader(input))
	sc.Scan()
	v_e := strings.Split(sc.Text(), " ")
	v, _ := strconv.Atoi(v_e[0])
	e, _ := strconv.Atoi(v_e[1])
	parents := make([]int, v+1)
	for i := range parents {
		parents[i] = i
	}
	
	var find_parent func(int) int
	
	find_parent = func(x int) int {
		if parents[x] != x {
			parents[x] = find_parent(parents[x])
		}
		return parents[x]
	}

	union_parent := func(a, b int) {
		if find_parent(a) > find_parent(b) {
			parents[a] = parents[b]
		} else {
			parents[b] = parents[a]			
		}
	}

	edges := make([]BuildCost, 0)
	result := 0

	for i:=0; i<e; i++ {
		abc := generateSeqFromStdin(sc, 3)
		buildCost := BuildCost{
			abc[2],
			Path{abc[0], abc[1]},
		}
		edges = append(edges, buildCost)
	}

	sort.SliceStable(edges, func(i, j int) bool{
		return edges[i].c < edges[j].c
	})
	
	for _, buildCost := range edges {
		a := buildCost.p.x
		b := buildCost.p.y
		if find_parent(a) == find_parent(b) {
			continue
		}
		union_parent(a, b)
		result += buildCost.c
	}

	fmt.Println(result)
}
