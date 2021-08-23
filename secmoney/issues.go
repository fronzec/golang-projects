package main

import (
	"fmt"
	"math/big"
)

func issues1() {
	var n float64 = 0
	for i := 0; i < 1000; i++ {
		n += .01
	}
	fmt.Println(n)
	var a float64 = .333
	fmt.Println(a+a+a)
}

func issues2()  {
	z, _ := new(big.Rat).SetString("1")
	three, _ := new(big.Rat).SetString("3")
	x := new(big.Rat).Quo(z, three)
	y := new(big.Rat).Quo(z, three)

	z = z.Sub(z, x)
	z = z.Sub(z, y)

	s := new(big.Rat).Add(x, y)
	s.Add(s, z)

	fmt.Println(x.FloatString(3), "+") // 0.333
	fmt.Println(y.FloatString(3), "+") // 0.333
	fmt.Println(z.FloatString(3))      // 0.333
	fmt.Println("=", s.FloatString(3)) // where did the other 0.001 go?
}

func issues3() {
	a := big.NewInt(1)
	b := big.NewInt(2)
	c := a // 1
	d := b // 2
	z := a.Add(a, b) // someone might misuse the API
	fmt.Println(a, b, z)
	fmt.Println(c.Cmp(d) < 0) // they might expect this to print "true"
	fmt.Println(c, d) // but instead, c was changed because it points to the same big.Int as a
}
