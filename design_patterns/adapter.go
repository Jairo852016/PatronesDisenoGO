package main

import "fmt"

type Payment interface {
	Pay()
}
type CashPayment struct {
}

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}
func PrcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct {
}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using Bankaccount %d \n", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	PrcessPayment(cash)
	/* bank := &BankPayment{}
	PrcessPayment(bank) */
	bpa := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}
	PrcessPayment(bpa)

}
