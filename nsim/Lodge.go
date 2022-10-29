package nsim

const lodgeBaseWood = 100
const lodgeBaseCost = 10

type Lodge struct {
	wood int
	cost int
	// wood cost?
}

func LodgeInit() Lodge {
	return Lodge{
		wood: lodgeBaseWood,
		cost: lodgeBaseCost,
	}
}

func LodgeWoodMod(l *Lodge, amount int) { // stores or consumes wood
	l.wood += amount
}

func LodgeWood(l *Lodge) int {
	return l.wood
}

func LodgeCost(l *Lodge) int {
	return l.cost
}
