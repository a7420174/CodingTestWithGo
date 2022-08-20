package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	[문제]

	n x m 크기의 금광이 있습니다. 금광은 1 x 1 크기의 칸으로 나누어져 있으며, 각 칸은 특정한 크기의 금이 들어 있습니다.
	
	채굴자는 첫 번째 열부터 출발하여 금을 캐기 시작합니다. 맨 처음에는 첫 번째 열의 어느 행에서든 출발할 수 있습니다.
	
	이후에 m - 1번에 걸쳐서 매번 오른쪽 위, 오른쪽, 오른쪽 아래 3가지 중 하나의 위치로 이동해야 합니다.
	
	결과적으로 채굴자가 얻을 수 있는 금의 최대 크기를 출력하는 프로그램을 작성하세요.
	
	[입력 조건]

	1. 첫째 줄에 테스트 케이스 T가 입력됩니다. (1 <= T <= 1000)
	
	2. 매 테스트 케이스 첫째 줄에 n과 m이 공백으로 구분되어 입력됩니다. (1 <= n, m <= 20)
	
	둘째 줄에 n x m개의 위치에 매장된 금의 개수가 공백으로 구분되어 입력됩니다. (1 <= 각 위치에 매장된 금의 개수 <= 100)
	
	[출력 조건]

	테스트 케이스마다 채굴자가 얻을 수 있는 금의 최대 크기를 출력합니다. 각 테스트 케이스는 줄 바꿈을 이용해 구분합니다.
	 */
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	T, _ := strconv.Atoi(sc.Text())

	for t:=0; t<T; t++ {
		sc.Scan()
		n_m := strings.Split(sc.Text(), " ")
		n, _ := strconv.Atoi(n_m[0])
		m, _ := strconv.Atoi(n_m[1])
		golds := generateSeqFromStdin(sc, n*m)
		var dp [3][4]int
		for i:=0;i<n;i++ {
			for j:=0;j<m;j++ {
				dp[i][j] = golds[i*m+j]
			}
		}
		for j:=1;j<m;j++ {
			for i:=0;i<n;i++ {
				var left_top, left, left_down int
				if i != 0 {
					left_top = dp[i-1][j-1]
				}
				if i != n-1 {
					left_down = dp[i+1][j-1]
				}
				left = dp[i][j-1]
				
				dp[i][j] += max(max(left_top, left), left_down)
			}
		}
		result := 0
		for i:=0;i<n;i++ {
			result = max(result, dp[i][m-1])
		}
		fmt.Println(dp)
		fmt.Println(result)
	}


  
}
