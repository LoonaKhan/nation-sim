package ppl

import (
	"nsim/utils/json"
)

// structs to structure job data

type ResourceProduction struct {
	BaseValue int  `json:"base-value"`
	Scalable  bool `json:"scalable"`
}

type Job struct {
	Income         ResourceProduction `json:"income"`
	FoodProduction ResourceProduction `json:"food-production"`
	WoodProduction ResourceProduction `json:"wood-production"`
	HappinessBonus int                `json:"happiness-bonus"`
	Name           string             `json:"name"`
}

var Jobs = json.Read[map[string]Job]("nsim/ppl/jobs.json")
