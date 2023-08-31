package context_test

import (
	"fmt"
	"sort"
	"sync"
	"testing"
)

func TestSliceAndArray(t *testing.T) {
	array := [5]int{0, 1, 2, 3, 4}
	slice := make([]int, 5)
	for i := 0; i < 5; i++ {
		slice[i] = i
	}
	newArray := array
	newArray[0] = 1
	newSlice := slice
	newSlice[0] = 1
	t.Log(array)
	t.Log(slice)
	t.Log(cap(slice))
	slice = append(slice, 5)
	t.Log(cap(slice))
	t.Log(cap(array))
	arr := []int{5, 4, 3, 2, 1}
	sort.Ints(arr)
	t.Log(arr)
}

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
