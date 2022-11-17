package main

import (
	"errors"
)

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var myStack Stack

	for i := 0; i < len(temperatures); i++ {
		for {
			curr, err := myStack.Top()
			if err != nil {
				myStack.Push(&Temp{
					index:       i,
					temperature: temperatures[i],
				})
				break
			}
			current := curr.(*Temp)
			if temperatures[i] > current.temperature {
				ans[current.index] = i - current.index
				myStack.Pop()
			} else {
				myStack.Push(&Temp{
					index:       i,
					temperature: temperatures[i],
				})
				break
			}
		}
	}
	return ans
}

type Stack []interface{}

type Temp struct {
	index       int
	temperature int
}

func (s *Stack) Push(num interface{}) {
	*s = append(*s, num)
}

func (s *Stack) Pop() (interface{}, error) {
	theStack := *s
	if len(*s) == 0 {
		return nil, errors.New("stackType is empty")
	}
	value := theStack[len(*s)-1]
	*s = theStack[:len(*s)-1]
	return value, nil
}

func (s *Stack) Top() (interface{}, error) {
	theStack := *s
	if len(theStack) == 0 {
		return nil, errors.New("stackType is empty")
	}
	return theStack[len(theStack)-1], nil
}
