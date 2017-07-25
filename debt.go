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
	DebtName           string
	DebtSize           int
	Currency           string
	Interest           float32
	Amortization       int
	QuantityOfPayments int
}

// NewDebt : returns a debt type.
func NewDebt(debtname, currency string, debtSize, payment int, interest float32) Debt {

	debt := Debt{}
	debt.DebtName = debtname
	debt.DebtSize = debtSize
	debt.Amortization = payment
	debt.Currency = currency
	debt.Interest = interest
	debt.QuantityOfPayments = initQuantityOfPayments(debtSize, payment)

	return debt
}

func initQuantityOfPayments(debtSize, amortization int) int {
	quantity := debtSize / amortization
	return quantity
}
