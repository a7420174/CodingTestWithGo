package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {
	/*
	출처: 이것이 취업을 위한 코딩 테스트다 with 파이썬 (나동빈 저)
	
	[문제]

	오늘 동빈이는 여행 가신 부모님을 대신해서 떡집 일을 하기로 했습니다. 오늘은 떡볶이 떡을 만드는 날입니다.
	
	동빈이네 떡볶이 떡은 재밌게도 떡볶이 떡의 길이가 일정하지 않습니다. 대신에 한 봉지 안에 들어가는 떡의 총 길이는
	
	절단기로 잘라서 맞춰줍니다.
	
	절단기에 높이(H)를 지정하면 줄지어진 떡을 한 번에 절단합니다. 높이가 H보다 긴 떡은 H 위의 부분이 잘릴 것이고,
	
	낮은 떡은 잘리지 않습니다. 예를 들어 높이가 19, 14, 10, 17cm 인 떡이 나란히 있고 절단기 높이를 15cm로 지정하면
	
	자른 뒤 떡의 높이는 15, 14, 10, 15cm가 될 것입니다. 잘린 떡의 길이는 차례대로 4, 0, 0, 2cm입니다.
	
	손님은 6cm만큼의 길이를 가져갑니다.
	
	손님이 왔을 때 요청한 총 길이가 M일 때 적어도 M만큼의 떡을 얻기 위해 절단기에 설정할 수 있는 높이의
	
	최댓값을 구하는 프로그램을 작성하세요.
	
	[입력 조건]

	1. 첫째 줄에 떡의 개수 N과 요청한 떡의 길이 M이 주어집니다. (1 <= N <= 1,000,000, 1 <= M <= 2,000,000,000)
	
	2. 둘째 줄에는 떡의 개별 높이가 주어집니다. 떡 높이의 총합은 항상 M 이상이므로, 손님은 필요한 양만큼
	
			떡을 사갈 수 있습니다. 높이는 10억보다 작거나 같은 양의 정수 또는 0 입니다.
	*/
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n_m := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(n_m[0])
	m, _ := strconv.Atoi(n_m[1])

	Hs := make([]int, n)
	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
	 Hs[i], _ = strconv.Atoi(s)
	}
	
	startTime := time.Now()

	sort.Slice(Hs, func(i, j int) bool {
		return Hs[i] < Hs[j]
	})

	var bisearch func(int, int) int
	bisearch = func(l, u int) int {
		if l > u {
			return u
		}
		mid := (l+u)/2
		total := 0
		for _, h := range Hs {
			if h > mid {
				total += h - mid
			}
		}
		switch {
			case total == m:
				return mid
			case total > m:
				return bisearch(mid+1, u)
			default:
				return bisearch(l, mid-1)
		}
	}

	result := bisearch(0, Hs[n-1])

	fmt.Println(result)
	
	endTime := time.Now()

	fmt.Println(endTime.Sub(startTime))
}
