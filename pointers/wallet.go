package main

type Wallet struct {
	accountBalance int
}

func (w *Wallet) Deposite(amount int) {
	w.accountBalance += amount
}

func (w *Wallet) Balance() int {
	return w.accountBalance
}
