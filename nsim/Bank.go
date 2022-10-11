package nsim

const bankBaseCost = 10
const bankBaseMoney = 100

type Bank struct {
	money int
	cost  int
}

func BankCon() Bank {
	return Bank{
		money: bankBaseMoney,
		cost:  bankBaseCost,
	}
}

func BankTransaction(b *Bank, amount int) { // works for both deposits and withdrawals
	b.money += amount
}

func BankCost(b *Bank) int { // getter for cost
	return b.cost
}

func BankMoney(b *Bank) int { // getter for money
	return b.money
}
