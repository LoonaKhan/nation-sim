package main

import nsim "nsim/nsim"

func main() { // todo: need a way to concurrently simulate this
	c := nsim.Country{
		Name:       "England",
		Bank:       nsim.Bank{Money: 5},
		Factories:  []nsim.Factory{},
		Population: []nsim.Person{},
	}

	nsim.Print(&c)
}
