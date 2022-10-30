package nsim

import "fmt"

const factoryBaseCost = 10
const factoryBaseIncome = 20
const factoryWoodCost = 10

type Factory struct {
	baseCost   int
	level      int
	baseIncome int
	woodCost   int
}

func FactoryInit() Factory { // factory constructor
	return Factory{
		baseCost:   factoryBaseCost,
		level:      1,
		baseIncome: factoryBaseIncome,
		woodCost:   factoryWoodCost,
	}
}

func FactoryIncome(f *Factory) int { // income scales with level
	return f.baseIncome * f.level
}

func FactoryCost(f *Factory) int { // cost scales with level? todo: cost does not scale with level. u build more buildings and they get more profitable
	return f.baseCost * f.level
}

func FactoryWoodCost(f *Factory) int {
	return f.woodCost
}

func FactoryLevel(f *Factory) int {
	return f.level
}

func FactoryLevelUp(f *Factory) {
	f.level++
}

func FactoriesCost(fac *[]Factory) int { // calculates the total cost of all factories in a country
	fCost := 0
	for f := range *fac {
		fCost += FactoryCost(&(*fac)[f])
	}
	return fCost
}

func FactoriesIncome(fac *[]Factory) int { // calculates the total income of all factories in a country
	fIncome := 0
	for f := range *fac {
		fIncome += FactoryIncome(&(*fac)[f])
	}
	return fIncome
}

func FactoriesWoodCost(fac *[]Factory) int {
	fCost := 0
	for f := range *fac {
		fCost += FactoryWoodCost(&(*fac)[f])
	}
	return fCost
}

func BuildFactory(fac *[]Factory) {
	*fac = append(*fac, FactoryInit())
}

func FactoryString(f *Factory) string {
	income, cost := FactoryIncome(f), FactoryCost(f)
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
		str += " {" + FactoryString(&(*fac)[f]) + "}\n"
	}
	str += "]"
	return str
}
