package decisions

/*
decisions are carried out by a country based on their happiness and money excess
happiness is based on several factors such as wood, food, pride.
*/

/*
Main actions are ranked by their urgency. based on the magnitude of their excesses
the higher the urgency, the quicker it shall be taken care of.
*/

/*
main actions and factors: (only 1 per turn)
money deficiency
	-> try to build more factories if can afford and workers from an industry in excess
	-> reduce army
money excess
	-> increase trade
	->
wood deficiency
	-> add more wood producing villagers from a resource in excess
	-> reduce buildings if in money excess
	-> reduce army
wood excess
	-> increase trade
food deficiency
	-> add more food producing villagers from a resource in excess
	-> reduce army
food excess
	-> increase trade
pride low
	-> increase army from a resource in excess
*/

/*
Bonus actions and factors: (1+ per turn alongside main action)
employ an unemployed villager
level up a villager
*/
