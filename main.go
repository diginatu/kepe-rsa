package main

import (
	"fmt"
	"math/big"
)

func LCM(a, b *big.Int) *big.Int {
	c := &big.Int{}
	c.GCD(nil, nil, a, b)

	r := &big.Int{}
	r.Mul(a, b)
	r.Div(r, c)

	return r
}

func main() {
	p, q := big.NewInt(101), big.NewInt(100003)
	n := &big.Int{}
	n.Mul(p, q)

	ps, qs := &big.Int{}, &big.Int{}
	ps.Sub(p, big.NewInt(1))
	qs.Sub(q, big.NewInt(1))

	lambdaN := LCM(ps, qs)

	e := big.NewInt(11)
	d := &big.Int{}
	(&big.Int{}).GCD(d, nil, e, lambdaN)
	if d.Sign() < 0 {
		d.Add(d, lambdaN)
	}

	message := big.NewInt(2525)
	c := &big.Int{}
	c.Exp(message, e, n)

	fmt.Println("encrypted: ", c)

	decrypted := &big.Int{}
	decrypted.Exp(c, d, n)

	fmt.Println(decrypted)
}
