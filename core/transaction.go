package core

type Transaction struct {
	From   string
	To     string
	Amount int
}

func NewTransaction(from, to string, amount int) Transaction {
	return Transaction{From: from, To: to, Amount: amount}
}
