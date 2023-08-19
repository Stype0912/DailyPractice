package context_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestContext(t *testing.T) {
	chan1 := make(chan int, 1)
	chan2 := make(chan int, 1)
	chan3 := make(chan int, 1)
	counter := 0
	var wg sync.WaitGroup
	wg.Add(300)
	func1 := func() {
		for counter < 100 {
			select {
			case <-chan1:
				fmt.Println(1)
				chan2 <- 1
				wg.Done()
			}
		}
	}
	func2 := func() {
		for {
			select {
			case <-chan2:
				fmt.Println(2)
				chan3 <- 1
				wg.Done()
			}
		}
	}
	func3 := func() {
		for {
			select {
			case <-chan3:
				fmt.Println(3)
				chan1 <- 1
				wg.Done()
			}
		}
	}
	chan1 <- 1
	go func1()
	go func2()
	go func3()
	wg.Wait()
}

func sortFloat64(a []float64) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
