package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	/*
	출처: 이것이 취업을 위한 코딩 테스트다 with 파이썬 (나동빈 저)
	
	효율적인 화폐 구성

	문제

	N가지 종류의 화폐가 있다. 이 화폐들의 개수를 최소한으로 이용해서 그 가치의 합이 M원이 되도록 하려고 한다. 이때 각 화폐는 몇 개라도 사용할 수 있으며, 사용한 화폐의 구성은 같지만 순서만 다른 것은 같은 경우로 구분한다. 예를 들어 2원, 3원 단위의 화폐가 있을 때는 15원을 만들기 위해 3원을 5개 사용하는 것이 가장 최소한의 화폐 개수이다.

	입력

	첫째 줄에 N,M이 주어진다(1<= N <= 100, 1<= M <= 10,000)
	이후의 N개의 줄에는 각 화폐의 가치가 주어진다. 화폐의 가치는 10,000보다 작거나 같은 자연수이다.
	출력

	첫째 줄에 경우의 수 X를 출력한다.(불가능할 때는 -1을 출력한다)
	 */
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n_m := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(n_m[0])
	m, _ := strconv.Atoi(n_m[1])

	coins := make([]int, 0)
	for i := 0; i < n; i++ {
		sc.Scan()
		coin, _ := strconv.Atoi(sc.Text())
		coins = append(coins, coin)
	}
	
	startTime := time.Now()

	dp := make([]int, m+1)
	for i:=1; i<=m; i++ {
		dp[i] = 10001
	}

	min := func(x, y int) int {
		if x <= y {
			return x
		} else {
			return y
		}
	}
	
	for _, c := range coins {
		for i:=c; i<=m; i++ {
		if dp[i-c] != 10001 {
			dp[i] = min(dp[i], dp[i-c]+1)
		}
		}
	}

	

	fmt.Println(dp[m])
	
	endTime := time.Now()

	fmt.Println(endTime.Sub(startTime))
}
