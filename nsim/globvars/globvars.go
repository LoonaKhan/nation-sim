package globvars

import "nsim/utils/json"

// todo: maybe just use a plain go file instead of a json?

type TownHall struct { // townhall
	Base_money int `json:"base-money"`
	Base_wood  int `json:"base-wood"`
	Base_food  int `json:"base-food"`
	Base_level int `json:"base-level"`
	Cost       int `json:"cost"`
}

type Country struct { // country
	Base_happ         int     `json:"base-happiness"`
	Excess_limiter    float64 `json:"excess-limiter"`
	Excess_happ_cap   int     `json:"excess-happiness-cap"`
	Pride_upper_limit int     `json:"pride-upper-limit"`
	Pride_lower_limit int     `json:"pride-lower-limit"`
	Pride_mod         float64 `json:"pride-modifier"`
}

type Person struct {
	Base_level int `json:"base-level"`
}

type Globvars struct { // globvars
	Townhall TownHall `json:"town-hall"`
	Country  Country  `json:"country"`
	Person   Person   `json:"person"`
}

var Globs = json.Read[Globvars]("nsim/globvars/globvars.json")
