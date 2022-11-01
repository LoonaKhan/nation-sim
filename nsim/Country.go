package nsim

import (
	"fmt"
	"math"
	"nsim/nsim/names"
	"nsim/nsim/pop"
)

const baseHapp = 2
const excessLimiter = 5.0
const excessHappCap = 10
const prideUpperLimit = 5
const prideLowerLimit = -1
const prideModifier = 0.25

type Country struct { // todo: make these encapsulated
	Name       string
	Happiness  float64       // Happiness is used as a factor for decision making
	Bank       Bank          // bank stores money and the building itself has a cost
	Lodge      Lodge         // lodge stores wood. wood is used for all infrastructure
	Silo       Silo          // Silo's store food. food is used to feed villagers
	Factories  []Factory     // factories generate money
	Population []pop.Person  // Population supplies people who can take on jobs
	Army       []*pop.Person // a list of all people people in the Army. just a list of references
}

func CountryInit(name string, initPeople int) *Country { // constructor
	c := Country{
		Name:       name,
		Happiness:  baseHapp,
		Bank:       BankInit(),
		Lodge:      LodgeInit(),
		Silo:       SiloInit(),
		Factories:  []Factory{FactoryInit()},
		Population: []pop.Person{},
		Army:       []*pop.Person{},
	}

	for i := 0; i < initPeople; i++ { // we initialize the country with a certain number of people
		pop.NewPerson(&c.Population, names.ChoosePersonName())
	}

	return &c
}

// GETTERS/SETTERS
func CountryHappiness(c *Country) float64 { // getter for Happiness
	return c.Happiness
}

func ModHappiness(c *Country, delta float64) {
	// we only change Happiness by modifying it, not setting it
	c.Happiness += delta
}

func ArmySize(c *Country) int { // gets the size of the country's Army
	return len(c.Army)
}

func NewSoldier(c *Country) {
	/*
		looks for the first citizen that is not a soldier
		switches their profession
		appends their reference to army
	*/

	for p := range c.Population {
		if pop.GetJob(&c.Population[p]) != "soldier" {
			pop.AssignJob(&c.Population[p], "soldier")
			c.Army = append(c.Army, &c.Population[p])
			return
		}
	}
}

// remove soldier

// soldier death?

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
	var costs int
	costs += FactoriesCost(&c.Factories)
	costs += BankCost(&c.Bank)
	costs += LodgeCost(&c.Lodge)
	costs += SiloCost(&c.Silo)

	// incomes
	var income int
	income += pop.PopIncome(&c.Population)
	income += FactoriesIncome(&c.Factories)

	BankTransaction(&c.Bank, (income - costs))
}

func calcFoodProduction(c *Country) {
	SiloFoodMod(&c.Silo, pop.PopFoodProduction(&c.Population))
}

func calcWoodProduction(c *Country) {
	popProduction := pop.PopWoodProduction(&c.Population) // dont have a way of producing wood yet
	factoriesProduction := FactoriesWoodCost(&c.Factories)
	LodgeWoodMod(&c.Lodge, (popProduction - factoriesProduction))
}

func calcHappiness(c *Country) { // calculates and applies the modification
	/*
		Happiness is calculated based on the pride of the people and the excess of resources.
		Resource excess has a higher weight
		Pride does not need to be too high, but negatives severely lower Happiness
		todo: modify this once u have implemented food and wood
	*/

	// pride from the army
	pride := math.Max(
		math.Min(math.Log2(float64(ArmySize(c))+prideModifier), prideUpperLimit),
		prideLowerLimit,
	)
	// calculates the excess of each resource
	foodExcess := math.Min(
		(float64(pop.PopFoodProduction(&c.Population)) / float64(excessLimiter)),
		excessHappCap)
	woodExcess := math.Min(
		(float64(pop.PopWoodProduction(&c.Population)-FactoriesWoodCost(&c.Factories)) / float64(excessLimiter)),
		excessHappCap)
	// happiness bonuses from the population
	happBonus := pop.PopHappBonus(&c.Population)

	// puts them into an equation
	delta := foodExcess + pride + woodExcess + float64(happBonus)
	ModHappiness(c, (delta - c.Happiness))
}

func Simulate(c *Country) {
	/* Simulates a country.
	calculates the economy
	tries to manage resources
	calculates the Happiness/wellbeing of the nation
	tries to manage the wellbeing
	manages trade/wars with other nations
	*/

	// todo: test
	calcFoodProduction(c)
	calcWoodProduction(c)
	calcHappiness(c)
	calcEconomy(c)
}

// TOSTRING METHOD
func CountryString(c *Country) string {
	return fmt.Sprintf("Country name: %s\n"+
		"Bank balance: %d\n"+
		"Food: %d\n"+
		"Wood: %d\n"+
		"Happiness: %.2f\n"+
		"num of Factories: %d\n"+
		"Population: %d\n"+
		"Army Size: %d\n",
		c.Name,
		BankMoney(&c.Bank),
		SiloFood(&c.Silo),
		LodgeWood(&c.Lodge),
		CountryHappiness(c),
		len(c.Factories),
		len(c.Population),
		len(c.Army),
	)
}
