package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Vertex struct{
	x, y int
}

func main() {
	/*
	
	 */
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n_m := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(n_m[0])
	m, _ := strconv.Atoi(n_m[1])

	graph := make([][]int, 0)
	for i := 0; i < n; i++ {
		sc.Scan()
		row := make([]int, 0)
		for _, r := range sc.Text() {
			v := int(r-'0')
			row = append(row, v)
		}
		graph = append(graph, row)
	}
	
	startTime := time.Now()

	moves := []Vertex{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	bfs := func(start Vertex) int {
		q := NewQueue()
		q.Push(start)
		for q.v.Len() > 0{
			v := q.Pop()
			for _, move := range moves {
				nx := v.(Vertex).x + move.x
				ny := v.(Vertex).y + move.y
				if nx<1||ny<1||nx>m||ny>n||(nx==1&&ny==1) {
					continue
				}
				if graph[ny-1][nx-1] == 0 {
					continue
				}
				if graph[ny-1][nx-1] == 1 {
					graph[ny-1][nx-1] = graph[v.(Vertex).y-1][v.(Vertex).x-1] + 1
					q.Push(Vertex{nx, ny})
				}
			}
		}
		return graph[n-1][m-1]
	}

	fmt.Println(bfs(Vertex{1,1}))
	
	endTime := time.Now()

	fmt.Println(endTime.Sub(startTime))
}
