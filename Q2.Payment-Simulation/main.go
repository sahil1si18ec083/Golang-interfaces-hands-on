package main

import "fmt"

type Payment interface {
	Pay(amount int)
}
type UPI struct {
}
type Card struct {
}

func (u UPI) Pay(amount int) {
	fmt.Printf("Paid ₹ %d using UPI", amount)
}
func (c Card) Pay(amount int) {
	fmt.Printf("Paid ₹ %d using Card", amount)
}
func Checkout(p Payment, amount int) {
	p.Pay(amount)

}
func main() {
	u := UPI{}
	c := Card{}
	Checkout(u, 123)
	Checkout(c, 1523)

}
