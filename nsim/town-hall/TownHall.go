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
		money: globvars.Globs.Townhall.Base_money,
		food:  globvars.Globs.Townhall.Base_food,
		wood:  globvars.Globs.Townhall.Base_wood,
		level: globvars.Globs.Townhall.Base_level,
		cost:  globvars.Globs.Townhall.Cost,
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
func Cost(t *TownHall) int {
	return t.cost
}
