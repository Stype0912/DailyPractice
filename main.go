package main

import (
	"Learning/zkrp/util"
	"fmt"
	"math/big"
)

func main() {
	a, n := big.NewInt(2), big.NewInt(7)
	fmt.Println(util.Inverse(a, n))
	a, b, n := big.NewInt(123), big.NewInt(777), big.NewInt(999)
	fmt.Println(util.PowerMod(a, b, n))
	//n = new(big.Int)
	n, _ = n.SetString("68312459744043142125516091584237965936489853199563456268379", 10)
	fmt.Println(util.IsPrime(n, 10))
}
