package nsim

import "fmt"

const personBaseCost = 10
const personBaseIncome = 1

type Person struct {
	Name       string
	Level      int
	baseCost   int
	baseIncome int
}

func PersonCon(name string) Person {
	/* person constructor.
	The reason this is a seperate function is because baseCost and baseIncome act as class variables
	*/
	return Person{
		Name:       name,
		Level:      1,
		baseCost:   personBaseCost,
		baseIncome: personBaseIncome,
	}
}

func PersonCost(p *Person) int { // should costs scale with level?
	return p.baseCost * p.Level
}

func PersonIncome(p *Person) int {
	return p.baseIncome * p.Level
}

func PersonLevelUp(p *Person) {
	p.Level++
}

func PopulationCost(pop *[]Person) int { // calculates the cost of all people in a country
	popCost := 0
	for p := range *pop {
		popCost += PersonCost(&(*pop)[p])
	}
	return popCost
}

func PopulationIncome(pop *[]Person) int { // calculates the income of all people in a country
	popIncome := 0
	for p := range *pop {
		popIncome += PersonIncome(&(*pop)[p])
	}
	return popIncome
}

func NewPerson(pop *[]Person, name string) {
	*pop = append(*pop, PersonCon(name))
}

func PersonString(p *Person) string {
	income, cost := PersonIncome(p), PersonCost(p)
	net := income - cost

	return fmt.Sprintf("Name: '%s', "+
		"Level: %d, "+
		"Cost: %d, "+
		"Income: %d, "+
		"Net_Profit: %d",
		p.Name, p.Level, cost, income, net)
}

func PopulationString(pop *[]Person) string {
	str := "Population: [\n"
	for p := range *pop {
		str += " {" + PersonString(&(*pop)[p]) + "}\n"
	}
	str += "]"
	return str
}
