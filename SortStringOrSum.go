package main

import (
	"fmt"
	"strings"
	"strconv"
	"unicode"
	"sort"
)

func main() {
	/*
	출처: 이것이 코딩 테스트다 with 파이썬 - 나동빈 저
	
	알파벳 대문자와 숫자(0~9)로만 구성된 문자열이 입력으로 주어집니다. 이때 모든 알파벳을 오름차순으로 정렬하여 이어서 출력한 뒤에, 그 뒤에 모든 숫자를 더한 값을 이어서 출력합니다.

	예를 들어 K1KA5CB7 이라는 값이 들어오면 ABCKK13을 출력합니다.

	(입력 조건)
	첫째 줄에 하나의 문자열 S가 주어집니다. (1 <= S의 길이 <= 10,000)
	
	(출력 조건)
	첫째 줄에 문제에서 요구하는 정답을 출력합니다.
	*/
	input := "AJKDLSI412K4JSJ9D"
	
	var result string
	var sum, num_cnt int
	
	for _, r := range input {
		if unicode.IsLetter(r) {
			result += string(r)
		} else {
			num, _ := strconv.Atoi(string(r))
			sum += num
			num_cnt++
		}
	}
	
	result = SortString(result)
	
	if num_cnt > 0 {
		result = result + strconv.Itoa(sum)
	}

	fmt.Println(result)

}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

