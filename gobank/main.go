package main

import (
	"cursos/gobank/accounts"
	"cursos/gobank/clients"
	"fmt"
)

func payBill(account VerifyAccount, billAmount float64) string {
	return account.Pay(billAmount)
}

type VerifyAccount interface {
	Pay(billAmount float64) string
}

func main() {
	willAccount := accounts.CurrentAccount{
		Holder:        clients.Holder{Name: "Will Smith", CPF: "123.456.789-00"},
		AgencyNumber:  123,
		AccountNumber: 456789,
	}

	joanAccount := accounts.CurrentAccount{
		Holder:        clients.Holder{Name: "Joan Doe", CPF: "987.654.321-00"},
		AgencyNumber:  321,
		AccountNumber: 987654,
	}

	willAccount.DepositAmount(1000)
	fmt.Println(willAccount.AmountToWithdraw(500))
	fmt.Println(willAccount.TransferAmount(200, &joanAccount))

	message, balance := willAccount.GetBalance()
	fmt.Println(message, balance)

	messageJoan, balanceJoan := joanAccount.GetBalance()
	fmt.Println(messageJoan, balanceJoan)

	johnAccount := accounts.SavingsAccount{
		Holder:        clients.Holder{Name: "John Doe", CPF: "111.222.333-44"},
		AgencyNumber:  555,
		AccountNumber: 666777,
	}
	janeAccount := accounts.SavingsAccount{
		Holder:        clients.Holder{Name: "Jane Roe", CPF: "444.333.222-11"},
		AgencyNumber:  888,
		AccountNumber: 999000,
	}

	johnAccount.DepositAmount(2000)
	fmt.Println(johnAccount.AmountToWithdraw(800))
	fmt.Println(johnAccount.TransferAmount(500, &janeAccount))
	messageJohn, balanceJohn := johnAccount.GetBalance()
	fmt.Println(messageJohn, balanceJohn)
	messageJane, balanceJane := janeAccount.GetBalance()
	fmt.Println(messageJane, balanceJane)

	fmt.Println(payBill(&willAccount, 100))
	fmt.Println(payBill(&johnAccount, 300))

}
