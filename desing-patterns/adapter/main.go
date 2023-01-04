package main

import "fmt"

type Payment interface {
	Pay()
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash")
}

type BankPayment struct{}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Pay using bank account: %d\n", bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	bpa := &BankPaymentAdapter{
		bankAccount: 777,
		BankPayment: &BankPayment{},
	}

	ProcessPayment(bpa)
}
