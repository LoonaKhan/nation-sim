package nsim

type Factory struct {
	baseCost   int
	level      int
	baseIncome int
}

func FactoryIncome(f *Factory) int { // income scales with level
	return f.baseIncome * f.level
}

func FactoryCost(f *Factory) int { // cost scales with level?
	return f.baseCost * f.level
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
