package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	a, ok := <-ch
	fmt.Println(a, ok)
}
