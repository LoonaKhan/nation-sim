package town_hall

import "nsim/nsim/globvars"

type TownHall struct {
	money int
	food  int
	wood  int
	level int
	cost  int // should it be free?
}

func Init() TownHall {
	return TownHall{
		money: globvars.ThGlob.BaseMoney,
		food:  globvars.ThGlob.BaseFood,
		wood:  globvars.ThGlob.BaseWood,
		level: globvars.ThGlob.BaseLevel,
		cost:  globvars.ThGlob.Cost,
	}
}

// Food
func Food(t *TownHall) int {
	return t.food
}
func FoodMod(t *TownHall, amount int) { // modifies the amount of food by a given value
	t.food += amount
}

// Wood
func Wood(t *TownHall) int {
	return t.wood
}
func WoodMod(t *TownHall, amount int) {
	t.wood += amount
}

// Money
func Money(t *TownHall) int {
	return t.money
}
func Transaction(t *TownHall, amount int) {
	t.money += amount
}

// Cost
func Cost(t *TownHall) int { // mantenance cost
	return t.cost
}
