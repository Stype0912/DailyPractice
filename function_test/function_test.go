package function_test

import (
	"sync"
	"testing"
	"time"
)

func Sum1(n int) int {
	ans := 0
	var wg sync.WaitGroup
	for i := 0; i <= n; i++ {
		wg.Add(1)
		go func(num int) {
			ans += num
			wg.Done()
		}(i)
	}
	wg.Wait()
	return ans
}

func Sum2(n int) int {
	ans := 0
	for i := 0; i <= n; i++ {
		ans += i
	}
	return ans
}

func TestFunc(t *testing.T) {
	timeStamp := time.Now()
	t.Log(Sum1(100))
	t.Log(time.Since(timeStamp).Nanoseconds())
}

func TestFunc1(t *testing.T) {
	timeStamp := time.Now()
	t.Log(Sum2(100))
	t.Log(time.Since(timeStamp).Nanoseconds())
}
