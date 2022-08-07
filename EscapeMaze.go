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
	출처: 이것이 취업을 위한 코딩 테스트다 with 파이썬 (나동빈 저)
	
	미로 탈출
	동빈이는 N X M 크기의 직사각형 형태의 미로에 갇혀 있다. 미로에는 여러 마리의 괴물이 있어 이를 피해 탈출해야 한다. 동빈이의 위치는 (1, 1) 이고 미로의 출구는 (N, M)의 위치에 존재하며 한 번에 한 칸씩 이동할 수 있다. 이때 괴물이 있는 부분은 0으로, 괴물이 없는 부분은 1로 표시되어 있다. 미로는 반드시 탈출할 수 있는 형태로 제시된다. 이때 동빈이가 탈출하기 위해 움직여야 하는 최소 칸의 개수를 구하시오. 칸을 셀 때는 시작 칸과 마지막 칸을 모두 포함해서 계산한다.
	입력 조건
	첫째 줄에 두 정수 N, M (4 <=N, M <= 200)이 주어집니다. 다음 N개의 줄에는 각각 M개의 정수(0 혹은 1)로 미로의 정보가 주어진다. 각각의 수들은 공백 없이 붙어서 입력으로 제시된다. 또한 시작 칸과 마지막 칸은 항상 1이다.
	출력 조건
	첫째 줄에 최소 이동 칸의 개수를 출력한다.
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
