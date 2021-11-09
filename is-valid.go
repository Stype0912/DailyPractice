package main

import "strings"

var zhan []string
var item string

func push(zhan []string, item string) []string {
	return append(zhan, item)
}
func pop(zhan []string) (string, []string) {
	item := zhan[len(zhan)-1]
	return item, zhan[:len(zhan)-1]
}

func isValid(s string) bool {
	sList := strings.Split(s, "")
	for _, L := range sList {
		if L == "(" || L == "{" || L == "[" {
			zhan = push(zhan, L)
		} else if len(zhan) == 0 {
			return false
		}
		if L == ")" {
			item, zhan = pop(zhan)
			if item != "(" {
				return false
			}
		}
		if L == "}" {
			item, zhan = pop(zhan)
			if item != "{" {
				return false
			}
		}
		if L == "]" {
			item, zhan = pop(zhan)
			if item != "[" {
				return false
			}
		}
	}
	if len(zhan) != 0 {
		return false
	}
	return true
}
