package main

import "fmt"

func getHint(secret string, guess string) string {
	var x, y int
	MapSecret, MapGuess := make(map[string]int), make(map[string]int)
	for i := 0; i < len(secret); i++ {
		itemA, itemB := string(secret[i]), string(guess[i])
		MapSecret[itemA]++
		MapGuess[itemB]++
		if itemA == itemB {
			x++
		}
	}
	for i, _ := range MapSecret {
		if MapGuess[i] == 0 {
			continue
		}
		y += Min(MapSecret[i], MapGuess[i])
	}
	y -= x
	return fmt.Sprintf("%vA%vB", x, y)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
