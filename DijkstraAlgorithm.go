package main

import (
	"fmt"
	"container/heap"
)

const INF int = 1e9

var n, m int = 6, 10

var graph [][]Item = [][]Item{
	{},
	{{2, 2}, {3, 5}, {4, 1}},
	{{3, 3}, {4, 2}},
	{{2, 3}, {6, 5}},
	{{3, 3}, {5, 1}},
	{{3, 1}, {6, 2}},
	{},
}

type Item struct {
	node, dist int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func main() {
	/*
	출처: 이것이 취업을 위한 코딩 테스트다 with 파이썬 (나동빈 저)

	다익스트라 최단경로 알고리즘 (개선된 구현 방법: 우선순위 큐)
	 */
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = INF
	}
	pq := &PriorityQueue{}
	heap.Init(pq)

	dist[1] = 0
	pq.Push(&Item{1, 0})
	
	for pq.Len() > 0 {
		item := pq.Pop().(*Item)
		if dist[item.node] < item.dist {
			continue
		}
		for _, j := range graph[item.node] {
			new_dist := j.dist + dist[item.node]
			if new_dist < dist[j.node] {
				dist[j.node] = new_dist
				pq.Push(&Item{j.node, new_dist})
			}
		}
	}

	for i, val := range dist {
		if i>0 && i<=n {
			fmt.Println("node", i, ":", val)
		}
		
	}
 }
