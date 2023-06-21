package fHi6rV

import "testing"

func TestFlip(t *testing.T) {
	t.Log(flipChess([]string{".X.", ".O.", "XO."}))
}
