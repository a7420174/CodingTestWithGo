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

	// var DFS func(int)
	// DFS = func(v int){
	// 	visited[v] = true
	// 	fmt.Printf("%v ", v)
	// 	for _, i := range graph[v] {
	// 		if !visited[i] {
	// 			DFS(i)
	// 		}
	// 	}
	// }
	// DFS(1)
	
	DFS(&graph, 1, &visited)
	fmt.Println()
}

func DFS(graph *[9][]int, v int, visited *[9]bool){
	visited[v] = true
	fmt.Printf("%v ", v)
	for _, i := range graph[v] {
	if !visited[i] {
	  DFS(graph, i, visited)
	}
  }
}
