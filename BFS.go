package main

import (
	"fmt"
)

func main() {
	/*
	출처: 이것이 취업을 위한 코딩 테스트다 with 파이썬 (나동빈 저)

	BFS(너비 우선 탐색) 예제
	*/
	graph := [9][]int{
		{},
		{2, 3, 8},
		{1, 7},
		{1, 4, 5},
		{3, 5},
		{3, 4},
		{7},
		{2, 6, 8},
		{1, 7},
	}

	var visited [9]bool
	q := NewQueue()
  q.Push(1)
  
  BFS(&graph, &visited, q)
	fmt.Println()
}

func BFS(graph *[9][]int, visited *[9]bool, q *Queue) {
  v := q.Pop()
  if v == nil {
    return
  }
	fmt.Printf("%v ", v)
	for _, i := range graph[v.(int)] {
		if !visited[i] {
			q.Push(i)
			visited[i] = true
		}
	}
  BFS(graph, visited, q)
}
