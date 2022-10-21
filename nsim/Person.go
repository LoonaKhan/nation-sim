package nsim

import "fmt"

// job titles are appended to the beginning of person names.
var jobs = map[string]map[string]int{ // jobs can specialize people
	"unemployed": {"baseCost": 2, "baseIncome": 3},                        // default
	"farmer":     {"baseCost": 2, "baseIncome": 5},                        // 2x food/income
	"soldier":    {"baseCost": 4, "baseIncome": 3, "happinessFactor": 5},  // contributes to happiness(pride), but also fear
	"nitwit":     {"baseCost": 2, "baseIncome": 0, "happinessFactor": -1}, // no income and useless
	"lumberjack": {},                                                      // still dont know. but he gathers wood
}

type Person struct {
	name       string
	job        string
	level      int
	baseCost   int
	baseIncome int
}

// INIT METHODS
func PersonInit(name string) Person {
	/* person constructor.
	The reason this is a seperate function is because baseCost and baseIncome act as class variables.
	Initially set each person as unemployed.
	countries can later assign each person a job and change their income
	*/
	return Person{
		name:       name,
		level:      1,
		job:        "unemployed",
		baseCost:   jobs["unemployed"]["baseCost"],
		baseIncome: jobs["unemployed"]["baseIncome"],
	}
}

func NewPerson(pop *[]Person, name string) { // adds a new member to the population
	// differs from the constructor since this appends a person to a population
	*pop = append(*pop, PersonInit(name))
}

// GETTERS/SETTERS
func PersonCost(p *Person) int { // getter for a person's cost
	return p.baseCost * p.level
}

func PersonIncome(p *Person) int { // getter for the person's income
	return p.baseIncome * p.level
}

func PersonLevel(p *Person) int { // getter for the person's level
	return p.level
}

func PersonName(p *Person) string {
	return p.name
}

func PersonLevelUp(p *Person) { // levels up a person
	p.level++
}

func PersonJob(p *Person) string { // getter for a person's job
	return p.job
}

func AssignJob(p *Person, newJob string) { // assigns a new job
	//oldJob := p.job
	if _, ok := jobs[newJob]; ok { // if the given job was in the job map
		p.job = newJob
		p.baseIncome = jobs[newJob]["baseIncome"]
		p.baseCost = jobs[newJob]["baseCost"]
	}

}

// POPULATION BASED METHODS.
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

// TOSTRING METHODS
func PersonString(p *Person) string { // toString method
	income, cost := PersonIncome(p), PersonCost(p)
	net := income - cost

	return fmt.Sprintf("name: '%s', "+
		"level: %d, "+
		"Cost: %d, "+
		"Income: %d, "+
		"Net_Profit: %d",
		p.name, p.level, cost, income, net)
}

func PopulationString(pop *[]Person) string { // toString method for the entire population
	str := "Population: [\n"
	for p := range *pop {
		str += " {" + PersonString(&(*pop)[p]) + "}\n"
	}
	str += "]"
	return str
}
