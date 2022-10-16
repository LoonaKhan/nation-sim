package nsim

import "fmt"

const countryBaseHappiness = 50

type Country struct { // todo: make these encapsulated
	Name       string
	happiness  int
	Bank       Bank
	Factories  []Factory
	Population []Person
}

func CountryInit(name string, initPeople int) Country { // constructor
	c := Country{
		Name:       name,
		happiness:  countryBaseHappiness,
		Bank:       BankInit(),
		Factories:  []Factory{},
		Population: []Person{},
	}

	for i := 0; i < initPeople; i++ { // we initialize the country with a certain number of people
		NewPerson(&c.Population, "")
	}

	return c
}

// GETTERS/SETTERS
func CountryHappiness(c *Country) int { // getter for happiness
	return c.happiness
}

// SIMULATERS
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

func calcHappiness(c *Country) { // calculates and modifies a country's happiness

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

// TOSTRING METHOD
func CountryString(c *Country) string {
	return fmt.Sprintf("Country name: %s\n"+
		"Bank balance: %d\n"+
		"num of Factories: %d\n"+
		"Population: %d\n",
		c.Name, c.Bank.money, len(c.Factories), len(c.Population))
}
