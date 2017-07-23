package main

import (
	"debt"
	"fmt"
)

func main() {

	d := debt.NewDebt("houseLoan", "dollar", 150000, 15000, 0)

	fmt.Println("NameOfDebt:Debt:Currency:Interest:Amortization:QuantityOfPayments")
	fmt.Println(d)
}
