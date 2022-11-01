package nsim

import (
	"fmt"
	"math"
	"nsim/nsim/names"
	"nsim/nsim/ppl"
	town_hall "nsim/nsim/town-hall"
)

const baseHapp = 2
const excessLimiter = 5.0
const excessHappCap = 10
const prideUpperLimit = 5
const prideLowerLimit = -1
const prideModifier = 0.25

type Country struct { // todo: make these encapsulated
	Name       string
	Happiness  float64            // Happiness is used as a factor for decision making
	TownHall   town_hall.TownHall // the townhall keeps track of food, money and wood
	Factories  []Factory          // factories generate money
	Population []ppl.Person       // Population supplies people who can take on jobs
	Army       []*ppl.Person      // a list of all people people in the Army. just a list of references
}

func CountryInit(name string, initPeople int) *Country { // constructor
	c := Country{
		Name:       name,
		Happiness:  baseHapp,
		TownHall:   town_hall.Init(),
		Factories:  []Factory{FactoryInit()},
		Population: []ppl.Person{},
		Army:       []*ppl.Person{},
	}

	for i := 0; i < initPeople; i++ { // we initialize the country with a certain number of people
		ppl.NewPerson(&c.Population, names.ChoosePersonName())
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
		if ppl.GetJob(&c.Population[p]) != "soldier" {
			ppl.AssignJob(&c.Population[p], "soldier")
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
	costs += town_hall.Cost(&c.TownHall)

	// incomes
	var income int
	income += ppl.PopIncome(&c.Population)
	income += FactoriesIncome(&c.Factories)

	town_hall.Transaction(&c.TownHall, (income - costs))
}

func calcFoodProduction(c *Country) {
	town_hall.FoodMod(&c.TownHall, ppl.PopFoodProduction(&c.Population))
}

func calcWoodProduction(c *Country) {
	popProduction := ppl.PopWoodProduction(&c.Population)  // production from the population
	factoriesProduction := FactoriesWoodCost(&c.Factories) // cost. from buildings
	town_hall.WoodMod(&c.TownHall, (popProduction - factoriesProduction))
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
		(float64(ppl.PopFoodProduction(&c.Population)) / float64(excessLimiter)),
		excessHappCap)
	woodExcess := math.Min(
		(float64(ppl.PopWoodProduction(&c.Population)-FactoriesWoodCost(&c.Factories)) / float64(excessLimiter)),
		excessHappCap)
	// happiness bonuses from the population
	happBonus := ppl.PopHappBonus(&c.Population)

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
	// todo add production, costs for each resource
	return fmt.Sprintf("Country name: %s\n"+
		"Bank balance: %d\n"+ // money prod/cost
		"Food: %d\n"+ // food prod/cost
		"Wood: %d\n"+ // wood prod/cost
		"Happiness: %.2f\n"+ // happiness prod/cost
		"num of Factories: %d\n"+ // total production from factories
		"Population: %d\n"+ // jobs of ppl
		"Army Size: %d\n",
		c.Name,
		town_hall.Money(&c.TownHall),
		town_hall.Food(&c.TownHall),
		town_hall.Wood(&c.TownHall),
		CountryHappiness(c),
		len(c.Factories),
		len(c.Population),
		len(c.Army),
	)
}
