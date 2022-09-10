package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var input string =
`5 5
1 2 3 2 5
`

func main() {
	/*
	출처: 이것이 코딩 테스트다 with 파이썬 - 나동빈 저
	
	투 포인터 알고리즘
	특정한 합을 가지는 부분 연속 수열 찾기
	 */
	sc := bufio.NewScanner(strings.NewReader(input))
	sc.Scan()
	n_m := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(n_m[0])
	m, _ := strconv.Atoi(n_m[1])
	seq := generateSeqFromStdin(sc, 5)

	var intervalSum, end, cnt int

	for start:=0; start<n; start++ {
		for intervalSum < m && end < n {
			intervalSum+=seq[end]
			end++
		}
		if intervalSum == m {
			cnt++
		}
		intervalSum -= seq[start]
	}

	fmt.Println(cnt)
}
