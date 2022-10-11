package nsim

type Bank struct {
	Money int
	cost  int
}

func BankTransaction(b *Bank, amount int) { // works for both deposits and withdrawals
	b.Money += amount
}

func BankCost(b *Bank) int { // getter for cost
	return b.cost
}

func BankMoney(b *Bank) int { // getter for Money
	return b.Money
}
