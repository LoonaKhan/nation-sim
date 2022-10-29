package nsim

import (
	"fmt"
	"math"
)

const countryBaseHappiness = 2
const excessLimiter = 5.0
const foodHappinessCap = 10
const woodHappinessCap = 10

type Country struct { // todo: make these encapsulated
	Name       string
	happiness  float64   // happiness is used as a factor for decision making
	Bank       Bank      // bank stores money and the building itself has a cost
	Lodge      Lodge     // lodge stores wood. wood is used for all infrastructure
	Silo       Silo      // Silo's store food. food is used to feed villagers
	Factories  []Factory // factories generate money
	Population []Person  // Population supplies people who can take on jobs
	army       []*Person // a list of all people people in the army. just a list of references
}

func CountryInit(name string, initPeople int) *Country { // constructor
	c := Country{
		Name:       name,
		happiness:  countryBaseHappiness,
		Bank:       BankInit(),
		Lodge:      LodgeInit(),
		Silo:       SiloInit(),
		Factories:  []Factory{FactoryInit()},
		Population: []Person{},
	}

	for i := 0; i < initPeople; i++ { // we initialize the country with a certain number of people
		NewPerson(&c.Population, ChoosePersonName())
	}

	return &c
}

// GETTERS/SETTERS
func CountryHappiness(c *Country) float64 { // getter for happiness
	return c.happiness
}

func ModHappiness(c *Country, delta float64) {
	// we only change happiness by modifying it, not setting it
	c.happiness += delta
}

func ArmySize(c *Country) int { // gets the size of the country's army
	return len(c.army)
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

func calcFoodProduction(c *Country) {
	//subtracts production by the consumption and modifies the silo
	totalProduction := PopulationFoodIncome(&c.Population)
	totalCost := PopulationFoodCost(&c.Population)

	SiloFoodMod(&c.Silo, (totalProduction - totalCost))
}

func calcHappiness(c *Country) { // calculates and applies the modification
	/*
		Happiness is calculated based on the pride of the people and the excess of resources.
		Resource excess has a higher weight
		Pride does not need to be too high, but negatives severely lower happiness
		todo: modify this once u have implemented food and wood
	*/

	pride := math.Log2(float64(ArmySize(c)) + 0.25) // pride
	// calculates the excess of each resource
	foodExcess := math.Min(
		(float64(PopulationFoodIncome(&c.Population)-PopulationFoodCost(&c.Population)) / float64(excessLimiter)),
		foodHappinessCap)

	// puts them into an equation
	delta := foodExcess + pride
	ModHappiness(c, (delta - c.happiness))
}

func Simulate(c *Country) {
	/* Simulates a country.
	calculates the economy
	tries to manage resources
	calculates the happiness/wellbeing of the nation
	tries to manage the wellbeing
	manages trade/wars with other nations
	*/

	// todo: test
	calcFoodProduction(c)
	calcHappiness(c)
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
