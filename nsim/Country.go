package nsim

import (
	"fmt"
	"math"
	"nsim/nsim/globvars"
	"nsim/nsim/names"
	"nsim/nsim/ppl"
	town_hall "nsim/nsim/town-hall"
)

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
		TownHall:   town_hall.Init(),
		Happiness:  globvars.Globs.Country.Base_happ,
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

func ModHappiness(c *Country, newValue float64) {
	// we only change Happiness by modifying it, not setting it
	c.Happiness = newValue
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

func foodExcess(c *Country) float64 {
	return math.Min(
		(float64(ppl.PopFoodProduction(&c.Population)) / float64(globvars.Globs.Country.Excess_limiter)),
		globvars.Globs.Country.Excess_happ_cap)
}

func calcWoodProduction(c *Country) {
	popProduction := ppl.PopWoodProduction(&c.Population)  // production from the population
	factoriesProduction := FactoriesWoodCost(&c.Factories) // cost. from buildings
	town_hall.WoodMod(&c.TownHall, (popProduction - factoriesProduction))
}

func woodExcess(c *Country) float64 {
	return math.Min(
		(float64(ppl.PopWoodProduction(&c.Population)-FactoriesWoodCost(&c.Factories)) / float64(globvars.Globs.Country.Excess_limiter)),
		globvars.Globs.Country.Excess_happ_cap)
}

func calcPride(c *Country) float64 {
	return math.Max(
		math.Min(math.Log2(float64(ArmySize(c))+globvars.Globs.Country.Pride_mod), globvars.Globs.Country.Pride_upper_limit),
		globvars.Globs.Country.Pride_lower_limit,
	)
}

func calcHappiness(c *Country) { // calculates and applies the modification
	/*
		Happiness is calculated based on the pride of the people and the excess of resources.
		Resource excess has a higher weight
		Pride does not need to be too high, but negatives severely lower Happiness
	*/

	// happiness bonuses from the population
	happBonus := ppl.PopHappBonus(&c.Population)

	// happiness is based on foodand wood excesses, pride and happiness bonuses
	happ := foodExcess(c) + calcPride(c) + woodExcess(c) + float64(happBonus)
	ModHappiness(c, happ)
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
