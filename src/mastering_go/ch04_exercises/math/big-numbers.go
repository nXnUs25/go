package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

var (
	pf  = fmt.Printf
	pln = fmt.Println
)

const precision = 600

/*
Using the math/big standard Go package, write a Go program that calculates square roots with high precision
	– choose the algorithm on your own.
*/
func main() {
	num := []int64{}

	for i := 0; i < 100; i += rand.Intn(10) {
		num = append(num, int64(i))
	}
	pln(num)

	for _, x := range num {
		v, e := NewtonSqrt(x)
		// i := IndianSqrt(x)
		pf("Newton Method - Sqrt of %v = %-100f\n", x, v)
		// pf("Indian Method - Sqrt of %v = %-100f\n", x, i)
		pf("math.Sqrt: %-f\n", math.Sqrt(float64(x)))
		pf("Where error is: %-5e\n", e)
		pln()
	}
}

//Babylonian method
func NewtonSqrt(num int64) (*big.Float, *big.Float) {

	// from doc but is good to write it self for better understanding
	// x{n+1} = 1/2 * ( xn + (2 / xn) )

	steps := int(math.Log2(precision))

	of := new(big.Float).SetPrec(precision).SetInt64(num)
	half := new(big.Float).SetPrec(precision).SetFloat64(0.5)
	x := new(big.Float).SetPrec(precision).SetInt64(1)

	t := new(big.Float).SetPrec(precision).SetInt64(0)

	for i := 0; i <= steps; i++ {
		t.Quo(of, x)   // t = 2/xn
		t.Add(x, t)    // t = xn + (2/xn)
		x.Mul(half, t) // xn+1 = 0.5 * t
	}
	t.Mul(x, x) // t = x*x

	return x, t.Sub(of, t)
}

// Bakhshali method
func IndianSqrt(x int64) *big.Float {

	// perfect squer S
	pSq := new(big.Float).SetPrec(precision).SetInt64(0)

	N := new(big.Float).SetPrec(precision).SetInt64(0)
	t := new(big.Float).SetPrec(precision).SetInt64(0)

	for i := x; i < 0; i-- {
		for j := int64(1); j < i; j++ {
			bigJ := new(big.Float).SetPrec(precision).SetInt64(j)
			t.Mul(bigJ, bigJ)
			if t.Cmp(big.NewFloat(float64(i))) == 0 {
				pSq.SetInt64(i)
				N.SetInt64(j)
				break
			}
		}
		if pSq.Cmp(big.NewFloat(0)) == 1 {
			break
		}
	}
	// calculate d
	d := new(big.Float)
	s := new(big.Float).SetPrec(precision).SetInt64(x)
	d.Sub(s, pSq)
	// calculate P
	two := new(big.Float).SetPrec(precision).SetInt64(2)
	P := new(big.Float).SetPrec(precision).SetInt64(0)
	t2 := new(big.Float).SetPrec(precision).SetInt64(0)
	t2.Mul(two, N)
	P.Quo(d, t2)
	// calculate A
	A := new(big.Float).SetPrec(precision).SetInt64(0)
	A.Add(N, P)

	// calculate sqr(x) == A - ((P * P) / (2.0 * A))
	// temp var to get 2*A
	t3 := new(big.Float).SetPrec(precision).SetInt64(0)
	t3.Mul(two, A)
	// P*P
	t4 := new(big.Float).SetPrec(precision).SetInt64(0)
	t4.Mul(P, P)
	// (P * P) / (2.0 * A)
	// t3/t4
	t5 := new(big.Float).SetPrec(precision).SetInt64(0)
	t5.Quo(t3, t4)
	// sqrt of x
	sqrt := new(big.Float).SetPrec(precision).SetInt64(0)
	sqrt.Sub(A, t5)
	return sqrt
}
