package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func getHint(secret string, guess string) string {
	var count1, count2 = map[string]int64{
		"0": 0,
		"1": 0,
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
		"7": 0,
		"8": 0,
		"9": 0,
	}, map[string]int64{
		"0": 0,
		"1": 0,
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
		"7": 0,
		"8": 0,
		"9": 0,
	}
	var x, y int64
	secretList := strings.Split(secret, "")
	guessList := strings.Split(guess, "")
	for i, _ := range secretList {
		if secretList[i] == guessList[i] {
			x++
		}
		count1[secretList[i]]++
		count2[guessList[i]]++
	}
	for i := int64(0); i < 10; i++ {
		y += Max(count1[strconv.FormatInt(i, 64)], count2[strconv.FormatInt(i, 10)])
	}
	xStr := strconv.FormatInt(x, 10)
	yStr := strconv.FormatInt(y-x, 10)
	return xStr + "A" + yStr + "B"
}

func Max(x int64, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func lastRemaining(n int, m int) int {
	var nList []int
	var index = 0
	for i := 0; i < n; i++ {
		nList = append(nList, i)
	}
	for i := int(0); i < n-1; i++ {
		index = (m + index - 1) % (n - i)
		nList = append(nList[:index], nList[index+1:]...)
	}
	return nList[0]
}

func main() {
	//时间戳
	t := time.Now()
	fmt.Println(t.Weekday().String())

}
