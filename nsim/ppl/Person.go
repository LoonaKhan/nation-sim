package ppl

import "fmt"

type Person struct {
	name  string
	job   Job
	level int
}

// INIT METHODS
func PersonInit(name string) Person {
	/* pop constructor.
	The reason this is a seperate function is because baseCost and baseIncome act as class variables.
	Initially set each pop as unemployed.
	countries can later assign each pop a job and change their income
	*/
	return Person{
		name:  name,
		level: 1,
		job:   Jobs["unemployed"],
	}
}

func NewPerson(pop *[]Person, name string) {
	// adds a new member to the population
	// differs from the constructor since this appends a pop to a population
	*pop = append(*pop, PersonInit(name))
}

// GETTERS/SETTERS

func Income(p *Person) int { // gets the income
	if p.job.Income.Scalable {
		return p.job.Income.BaseValue * p.level
	}
	return p.job.Income.BaseValue
}

func FoodProduction(p *Person) int { // gets the food production
	if p.job.FoodProduction.Scalable {
		return p.job.FoodProduction.BaseValue * p.level
	}
	return p.job.FoodProduction.BaseValue
}

func WoodProduction(p *Person) int { // gets the wood production
	if p.job.WoodProduction.Scalable {
		return p.job.WoodProduction.BaseValue * p.level
	}
	return p.job.WoodProduction.BaseValue
}

func Level(p *Person) int { // gets for the pop's level
	return p.level
}

func Name(p *Person) string { // gets the name
	return p.name
}

func LevelUp(p *Person) { // levels up a pop
	p.level++
}

func GetJob(p *Person) string { // getter for a pop's job
	return p.job.Name
}

func AssignJob(p *Person, newJob string) { // assigns a new job
	if _, ok := Jobs[newJob]; ok { // if the given job was in the job map
		p.job = Jobs[newJob]
		p.level = 1 // resets level
	}
}

func happBonus(p *Person) int {
	return p.job.HappinessBonus
}

// POPULATION BASED METHODS.

func PopIncome(pop *[]Person) int { // calculates the income of all people in a country
	income := 0
	for p := range *pop {
		income += Income(&(*pop)[p])
	}
	return income
}

func PopFoodProduction(pop *[]Person) int {
	production := 0
	for p := range *pop {
		production += FoodProduction(&(*pop)[p])
	}
	return production
}

func PopWoodProduction(pop *[]Person) int {
	production := 0
	for p := range *pop {
		production += WoodProduction(&(*pop)[p])
	}
	return production
}

func PopHappBonus(pop *[]Person) int {
	bonus := 0
	for p := range *pop {
		bonus += happBonus(&(*pop)[p])
	}
	return bonus
}

// TOSTRING METHODS
func PersonString(p *Person) string { // toString method
	return fmt.Sprintf("name: '%s', "+
		"level: %d, "+
		"Job: %s\n",
		p.name, p.level, p.job.Name)
}

func PopulationString(pop *[]Person) string { // toString method for the entire population
	str := "Population: [\n"
	for p := range *pop {
		str += " {" + PersonString(&(*pop)[p]) + "}\n"
	}
	str += "]"
	return str
}
