package money

type Money struct {
	Amount   int
	Currency string
}

func New(amt int, cur string) Money {
	return Money{Amount: amt, Currency: cur}
}
