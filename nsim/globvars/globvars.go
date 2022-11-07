package globvars

import "nsim/utils/json"

// todo: maybe just use a plain go file instead of a json?

type TownHall struct { // townhall
	BaseMoney int `json:"base-money"`
	BaseWood  int `json:"base-wood"`
	BaseFood  int `json:"base-food"`
	BaseLevel int `json:"base-level"`
	Cost      int `json:"cost"`
}

type Country struct { // country
	BaseHapp      float64 `json:"base-happiness"`
	ExcessLim     float64 `json:"excess-limiter"`
	ExcessHappCap float64 `json:"excess-happiness-cap"`
	PrideUl       float64 `json:"pride-upper-limit"`
	PrideLl       float64 `json:"pride-lower-limit"`
	PrideMod      float64 `json:"pride-modifier"`
}

type Person struct {
	BaseLevel int `json:"base-level"`
}

type Fac struct {
	Cost           int `json:"cost"`
	BaseWoodCost   int `json:"base-wood-cost"`
	BaseProduction int `json:"base-wood-production"`
}

type Globvars struct { // globvars
	Townhall TownHall `json:"town-hall"`
	Country  Country  `json:"country"`
	Person   Person   `json:"person"`
	Factory  Fac      `json:"factory"`
}

var Globs = json.Read[Globvars]("nsim/globvars/globvars.json")
var FacGlob = Globs.Factory
var PplGlob = Globs.Person
var ThGlob = Globs.Townhall
var CGlob = Globs.Country
