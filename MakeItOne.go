package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strconv"
)

func main() {
	/*
	출처: 이것이 취업을 위한 코딩 테스트다 with 파이썬 (나동빈 저)
	
	1로 만들기

	문제

	정수 X가 주어질때 정수 X에 사용할 수 있는 연산은 다음과 같이 4가지이다.

	1) X가 5로 나누어떨어지면, 5로 나눈다.

	2) X가 3으로 나누어 떨어지면, 3으로 나눈다.

	3) X가 2로 나누어 떨어지면, 2로 나눈다.

	4) X에서 1을 뺀다.

	정수 X가 주어졌을때, 연산 4개를 적절히 사용해서 1을 만들어야한다. 이 연산을 사용하는 횟수의 최솟값을 출력해라.

	X = 26일 경우
	1. 26 - 1 = 25
	2. 25 /5 = 5
	3. 5 / 5 = 1

	입력

	첫째 줄에 정수 X이 주어진다. (1<=X<=30,000)
	출력

	첫째 줄에 연산을 하는 횟수의 최솟값을 출력한다.
	 */
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	startTime := time.Now()

	dp := make([]int, n+1)

	for i:=1; i<=n; i++ {
		cnt := dp[i-1]+1
		if i%5 == 0 && dp[i/5]+1 < cnt {
			cnt = dp[i/5]+1
		}
		if i%3 == 0 && dp[i/3]+1 < cnt {
			cnt = dp[i/3]+1
		}
		if i%2 == 0 && dp[i/2]+1 < cnt {
			cnt = dp[i/2]+1
		}
		dp[i] = cnt
	}

	fmt.Println(dp[n])
	
	endTime := time.Now()

	fmt.Println(endTime.Sub(startTime))
}
