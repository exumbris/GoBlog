package main

type Counter struct {
	count int
}

type BankAccount struct {
	balance float64
}

type Player struct {
	name string
	health int
}

func reduceHealth(myPlayer Player) {
	myPlayer.health -= 10
}

func (account *BankAccount) Deposit(amountToDeposit float64) {
	account.balance += amountToDeposit
}

func (account *BankAccount) Withdraw(amountToWithdraw float64) {
	account.balance -= amountToWithdraw
}

func (counter *Counter) Increment() {
	counter.count += 1
}

func DoubleScored(val *int) {
	*val *= 2
}



