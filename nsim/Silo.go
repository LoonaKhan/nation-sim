package nsim

const siloBaseFood = 100
const siloBaseCost = 10

type Silo struct {
	food int
	cost int
	// wood cost?
}

func SiloInit() Silo {
	return Silo{
		food: siloBaseFood,
		cost: siloBaseCost,
	}
}

func SiloFoodMod(s *Silo, amount int) { // stores or consumes food
	s.food += amount
}

func SiloFood(s *Silo) int {
	return s.food
}

func SiloCost(s *Silo) int {
	return s.cost
}
