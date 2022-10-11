package nsim

type Person struct {
	Name       string
	Level      int
	BaseCost   int
	BaseIncome int
}

func PersonCost(p *Person) int { // should costs scale with level?
	return p.BaseCost * p.Level
}

func PersonIncome(p *Person) int {
	return p.BaseIncome * p.Level
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
