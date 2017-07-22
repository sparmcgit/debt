/*
	Package debt describes a debt.
	Martin Häggström 170722
*/
package debt

// Debt : type
// debtName : name of the debt.
// debtSize : size of the debt.
// payPerMont : how much to pay per month.
// currency : currency.
// quantityOfMonths : quantity of months it takes to pay the debt.
type Debt struct {
	debtName         string
	debtSize      int
	payPerMonth    int
	currency            string
	interest	float32
	amortization int
	quantityOfMonths int
}