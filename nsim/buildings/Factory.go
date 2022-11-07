package buildings

import (
	"fmt"
	"nsim/nsim/globvars"
)

type Factory struct {
	baseCost   int
	level      int
	baseIncome int
	woodCost   int
}

func FactoryInit() Factory { // factory constructor
	return Factory{
		baseCost:   globvars.FacGlob.Cost,
		level:      1,
		baseIncome: globvars.FacGlob.BaseProduction,
		woodCost:   globvars.FacGlob.BaseWoodCost,
	}
}

func Income(f *Factory) int { // income scales with level
	return f.baseIncome * f.level
}

func Cost(f *Factory) int { // cost to initially build
	return f.baseCost
}

func WoodCost(f *Factory) int { // maintenance
	return f.woodCost
}

func Level(f *Factory) int {
	return f.level
}

func LevelUp(f *Factory) {
	f.level++
}

func FactoriesCost(fac *[]Factory) int { // calculates the total cost of all factories in a country
	fCost := 0
	for f := range *fac {
		fCost += Cost(&(*fac)[f])
	}
	return fCost
}

func FactoriesIncome(fac *[]Factory) int { // calculates the total income of all factories in a country
	fIncome := 0
	for f := range *fac {
		fIncome += Income(&(*fac)[f])
	}
	return fIncome
}

func FactoriesWoodCost(fac *[]Factory) int {
	fCost := 0
	for f := range *fac {
		fCost += WoodCost(&(*fac)[f])
	}
	return fCost
}

func Build(fac *[]Factory) {
	*fac = append(*fac, FactoryInit())
}

func ToString(f *Factory) string {
	income, cost := Income(f), Cost(f)
	net := income - cost
	return fmt.Sprintf("level: %d, "+
		"Cost: %d, "+
		"Income: %d, "+
		"Net_Profit: %d",
		f.level, cost, income, net)
}

func FactoriesString(fac *[]Factory) string {
	str := "Factories: [\n"
	for f := range *fac {
		str += " {" + ToString(&(*fac)[f]) + "}\n"
	}
	str += "]"
	return str
}
