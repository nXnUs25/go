// struct demo
package main

import (
	"fmt"
)

// Trade is a trade in stocks
type Trade struct {
	Symbol string  // Stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade price
	Buy    bool    // true if buy trade, false if sell trade
}

func main() {
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)

	fmt.Printf("%+v\n", t1)

	fmt.Println(t1.Symbol)

	t2 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)

	symbol := "HHH"
	vol := 2
	p := 22.2
	buy := false
	t4 := Trade{symbol, vol, p, buy}
	fmt.Printf("%+v\n", t4)

	symbol1 := "HHHLLL"
	vol1 := 10
	p1 := 23.2
	buy1 := false
	t5 := &Trade{symbol1, vol1, p1, buy1}
	fmt.Printf("%+v\n", t5)

	t6 := &Trade{
		Symbol: symbol,
		Volume: 1033,
		Price:  91.98,
		Buy:    true,
	}
	fmt.Printf("%+v\n", t6)

}
