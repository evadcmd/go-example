package main

import "fmt"

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("BTC: %d coins", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(b Bitcoin) {
	w.balance += b
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) String() string {
	return fmt.Sprintf("%s", w.balance)
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("name: %s,  age: %d", p.Name, p.Age)
}

func main() {
	w := Wallet{}
	w.Deposit(5) // Is is OK!!
	fmt.Println(w.Balance())

	p := Person{"TATA", 32}
	fmt.Println(p)  // {TATA 32}
	fmt.Println(&p) // name: TATA,  age: 32
}
