package nsim

import "fmt"

type Country struct {
	Name       string
	Bank       Bank
	Factories  []Factory
	Population []Person
}

func Print(c *Country) {
	fmt.Printf("Country Name: %s\n"+
		"Bank balance: %d\n"+
		"num of Factories: %d\n"+
		"Population: %d",
		c.Name, c.Bank.Money, len(c.Factories), len(c.Population))
}

func Simulate(c *Country) {
	/*
		simulates a country.

		Adds up all the costs of each component in the country;
			population
			factories
			banks

		and then adds up all of their incomes
	*/

	// costs
	var totalCosts int
	totalCosts += PopulationCost(&c.Population)
	totalCosts += FactoriesCost(&c.Factories)
	totalCosts += BankCost(&c.Bank)

	// incomes
	var totalIncome int
	totalIncome += PopulationIncome(&c.Population)
	totalIncome += FactoriesIncome(&c.Factories)

	BankTransaction(&c.Bank, (totalIncome - totalCosts))

}
