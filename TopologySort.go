package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var input string = 
`7 8
1 2
1 5
2 3
2 6
3 4
4 7
5 6
6 4
`

func main() {
	/*
	출처: 이것이 코딩 테스트다 with 파이썬 - 나동빈 저
	위상 정렬 알고리즘
	 */
	sc := bufio.NewScanner(strings.NewReader(input))
	sc.Scan()
	v_e := strings.Split(sc.Text(), " ")
	v, _ := strconv.Atoi(v_e[0])
	e, _ := strconv.Atoi(v_e[1])
	indegree := make([]int, v+1)
	graph := make([][]int, v+1)
	
	for i:=0; i<e; i++ {
		path := generateSeqFromStdin(sc, 2)
		graph[path[0]] = append(graph[path[0]], path[1])
		indegree[path[1]] += 1
	}

	q := NewQueue()
	
	for i, j := range indegree {
		if i>0 && j == 0 {
			q.Push(i)
		}
	}
	
	result := make([]int, 0)
	
	for q.v.Len() > 0 {
		now := q.Pop().(int)
		result = append(result, now)
		for _, node := range graph[now] {
			indegree[node] -= 1
			if indegree[node] == 0 {
				q.Push(node)
			}
		}
	}
}
