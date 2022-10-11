package main

import (
	"fmt"
	"nsim/nsim"
)

func main() { // todo: need a way to concurrently simulate this
	c := nsim.Country{
		Name:       "England",
		Bank:       nsim.BankCon(),
		Factories:  []nsim.Factory{},
		Population: []nsim.Person{},
	}

	fmt.Println(nsim.CountryString(&c)) // before

	// adds people, a factory and simulates the economy
	nsim.NewPerson(&c.Population, "person")
	nsim.BuildFactory(&c.Factories)
	nsim.Simulate(&c)

	// prints data afterward
	fmt.Println(nsim.CountryString(&c))
	fmt.Println(nsim.PopulationString(&c.Population))
	fmt.Println(nsim.FactoriesString(&c.Factories))

}
