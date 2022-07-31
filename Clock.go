package main

import (
	"fmt"
    "strconv"
    "strings"
)

func main() {
	/*
	출처: 이것이 취업을 위한 코딩 테스트다 with 파이썬 (나동빈 저)
	
	정수 N이 입력되면 00시 00분 00초부터 N시 59분 59초까지의 모든 시각 중에서 3이 하나라도 포함되는
	모든 경우의 수를 구하는 프로그램을 작성하라. 예를 들어 1을 입력했을 때
	다음은 3이 하나라도 포함되어 있으므로 세어야 하는 시각이다

	00시 00분 03초
	00시 13분 30초
	반면에 다음은 3이 하나도 포함되어 있지 않으므로 세면 안 되는 시각이다

	00시 02분 55초
	01시 27분 45초
	입력

	첫째 줄에 정수 N이 입력된다.(0<=N<=23)

	출력

	00시 00분 00초부터 N시 59분 59초까지의 모든 시각 중에서 3이 하나라도 포함되는 모든 경우의 수를 출력한다.

	입력 예시
	5

	출력 예시
	11475
	*/
	
	input := "5"
	N, _ := strconv.Atoi(input)
	
	var cnt int
	for i := 0; i <= N; i++ {
		for j := 0; j < 60; j++ {
			for k := 0; k < 60; k++ {
				if (strings.Contains(strconv.Itoa(i), "3") ||
					strings.Contains(strconv.Itoa(j), "3") ||
					strings.Contains(strconv.Itoa(k), "3")) {
					cnt++
				}
			}
		}
	}
	
	fmt.Println(cnt)
}
