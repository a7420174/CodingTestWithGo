package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	/*
	출처: 이것이 취업을 위한 코딩 테스트다 with 파이썬 (나동빈 저)

	N명의 병사가 무작위로 나열되어 있습니다. 각 병사는 특정한 값의 전투력을 보유하고 있습니다.
	
	병사를 배치할 때는 전투력이 높은 병사가 앞쪽에 오도록 내림차순으로 배치를 하고자 합니다. 
	
	다시 말해 앞쪽에 있는 병사의 전투력이 항상 뒤쪽에 있는 병사보다 높아야 합니다.
	
	또한 배치 과정에서는 특정한 위치에 있는 병사를 열외시키는 방법을 이용합니다. 그러면서도 남아 있는 병사의 수가
	
	최대가 되도록 하고 싶습니다. 예를 들어, N = 7일 때 나열된 병사들의 전투력이 다음과 같다고 가정하겠습니다.

	병사 번호	1	2	4	5	7
	전투력	15	11	8	5	4

 	병사에 대한 정보가 주어졌을 때, 남아있는 병사의 수가 최대가 되도록 하기 위해서 열외해야 하는 병사의 수를 출력하는 프로그램을 작성하세요.

	입력
	첫째 줄에 N이 주어진다. (1 ≤ N ≤ 2,000) 둘째 줄에 각 병사의 전투력이 공백을 기준으로 구분되어 차례대로 주어진다. 각 병사의 전투력은 10,000,000보다 작거나 같은 자연수이다.

	출력
	첫째 줄에 남아있는 병사의 수가 최대가 되도록 하기 위해서 열외해야 하는 병사의 수를 출력한다.
	 */
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	arr := generateSeqFromStdin(sc, n)

	dp := make([]int, n)
	dp[0] = 1
	for i:=1;i<n;i++ {
		for j:=0;j<i;j++ {
			if arr[j] > arr[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	max_len := 0
	for _, val := range dp {
		max_len = max(max_len, val)
	}
	fmt.Println(n - max_len)
 }
