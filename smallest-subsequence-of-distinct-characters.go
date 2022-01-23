package main

import (
	"errors"
	"strings"
)

type stackType []string

func (s *stackType) Push(new string) {
	newS := *s
	newS = append(newS, new)
}

func (s *stackType) Pop() (result string, err error) {
	newS := *s
	if len(newS) <= 0 {
		return "", errors.New("stackType is empty")
	}
	result = newS[len(newS)-1]
	newS = newS[:len(newS)-1]
	return
}

func (s *stackType) Top() (result string, err error) {
	newS := *s
	if len(newS) <= 0 {
		return "", errors.New("stackType is empty")
	}
	result = newS[len(newS)-1]
	return
}
func smallestSubsequence(s string) (result string) {
	//stack := new(stackType)
	HashMap := map[string]int{}
	HashBoolMap := map[string]bool{}
	for _, i := range strings.Split(s, "") {
		HashMap[i]++
	}

	for _, i := range strings.Split(s, "") {
		if HashMap[i] == 1 {
			result += i
			continue
		}
		if HashBoolMap[i] == true {

		}
	}
	return ""
}
