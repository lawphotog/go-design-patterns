package main

import "fmt"

type BankAccount struct {
	balance int
}

func NewBankAccount() *BankAccount {
	return &BankAccount{
		balance: 0,
	}
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount int) {
	b.balance -= amount
}

type BankAccontAction int

const (
	Deposit BankAccontAction = iota
	Withdraw
)

type BankAccountCommand struct {
	action  BankAccontAction
	account *BankAccount
}

func NewBankAccountCommand(ba *BankAccount) *BankAccountCommand {
	return &BankAccountCommand{
		account: ba,
	}
}

func (b *BankAccountCommand) Call(action BankAccontAction, amount int) {
	switch action {
	case Deposit:
		b.account.Deposit(amount)
	case Withdraw:
		b.account.Withdraw(amount)
	}
}

func main() {
	ba := NewBankAccount()
	command := NewBankAccountCommand(ba)
	command.Call(Deposit, 100)
	command.Call(Withdraw, 60)

	fmt.Println(ba.balance) //expect 40
}
