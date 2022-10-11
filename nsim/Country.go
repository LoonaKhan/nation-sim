package nsim

import "fmt"

type Country struct {
	Name       string
	Bank       Bank
	Factories  []Factory
	Population []Person
}

func CountryString(c *Country) string {
	return fmt.Sprintf("Country Name: %s\n"+
		"Bank balance: %d\n"+
		"num of Factories: %d\n"+
		"Population: %d\n",
		c.Name, c.Bank.money, len(c.Factories), len(c.Population))
}

func calcEconomy(c *Country) {
	/*
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

func Simulate(c *Country) {
	/* Simulates a country.
	calculates the economy
	tries to manage resources
	calculates the happiness/wellbeing of the nation
	tries to manage the wellbeing
	manages trade/wars with other nations
	*/

	calcEconomy(c)
}
