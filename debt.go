/*
	Copyright 2017 Martin Häggström <hejpadej@spray.se>.
	All rights reserved. Use of this source code is governed by a
	BSD two clause license that can be found in the LICENSE file.
*/

// Package debt describes a debt.
// Martin Häggström 170722
package debt

// Debt : type
// debtName : name of the debt.
// debtSize : size of the debt.
// amortization : how much to amortizise.
// currency : currency.
// quantityOfPayments : quantity of payments to pay the debt.
type Debt struct {
	debtName           string
	debtSize           int
	currency           string
	interest           float32
	amortization       int
	quantityOfPayments int
}

// InitDebt : returns a debt type.
func InitDebt(debtname, currency string, debtSize, payment int, interest float32) Debt {

	debt := Debt{}
	debt.debtName = debtname
	debt.debtSize = debtSize
	debt.amortization = payment
	debt.currency = currency
	debt.interest = interest
	debt.quantityOfPayments = initQuantityOfPayments(debtSize, payment)

	return debt
}

func initQuantityOfPayments(debtSize, amortization int) int {
	quantity := debtSize / amortization
	return quantity
}
