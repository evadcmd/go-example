package main

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(b Bitcoin) {
	w.balance += b
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {
	w := Wallet{}
	w.Deposit(5) // Is is OK!!
	fmt.Println(w.Balance())
}
