package main

import (
	"fmt"
    "strconv"
    "strings"
	"sort"
)

func main() {
	/*
	출처: 이것이 취업을 위한 코딩 테스트다 with 파이썬 (나동빈 저)

	- 한 마을에 모험가가 N명
	- 모험가별로 가지고 있는 ‘공포도’가 다르다.
	- 공포도가 X인 모험가는 반드시 X명 이상으로 구성한 모험가 그룹에 참여
	- 최대 몇 그룹이 탄생하는가
	*/
	input := "3 2 3 3 2 1 2 1 3"
	s_str := strings.Split(input, " ")
	var s_int []int
	for _, num := range s_str {
		score, _ := strconv.Atoi(num)
		s_int = append(s_int, score)
	}	
	
	sort.Slice(s_int, func(i, j int) bool {
		return s_int[i] < s_int[j]
	})
	
	var cnt, n_mem int
	for _, score := range s_int {
		n_mem++
		if n_mem >= score {
			cnt++
			n_mem=0
		}
	}
	
	fmt.Println(cnt)
}
