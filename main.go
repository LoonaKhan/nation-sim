package main

import (
	"fmt"
	"nsim/nsim"
	"sync"
)

var BUFFER_SIZE = 2
var TURNS = 2

func SimTurn(cc chan nsim.Country, turn int) chan nsim.Country {
	/*
		simulates a turn.
		all countries are simulated concurrently.
		transfers all countries to a new channel of results to be returned
	*/

	simulated := make(chan nsim.Country, BUFFER_SIZE)

	wg := sync.WaitGroup{}
	for g := 0; g < BUFFER_SIZE; g++ { // for all countries in the channel
		wg.Add(1)
		//g := g
		go func() { // takes in countries, modifies and returns
			defer wg.Done()
			c := <-cc
			nsim.Simulate(&c)
			fmt.Println(nsim.CountryString(&c))
			simulated <- c
		}()
	}
	close(cc)

	wg.Wait()
	fmt.Println("Done turn:", turn)

	return simulated

}

func Sim() { // simulates countries

	CountryChan := make(chan nsim.Country, BUFFER_SIZE)

	for c := 0; c < BUFFER_SIZE; c++ { // fills up the buffer beforehand
		CountryChan <- nsim.Country{
			Name: fmt.Sprintf("Country %d", c),
			Bank: nsim.BankInit(),
			Factories: []nsim.Factory{
				nsim.FactoryInit(),
			},
			Population: []nsim.Person{
				nsim.PersonInit(""),
				nsim.PersonInit(""),
				nsim.PersonInit(""),
				nsim.PersonInit(""),
				nsim.PersonInit(""),
			},
		}
	}

	for i := 0; i < TURNS; i++ { // each iteration is its own turn
		CountryChan = SimTurn(CountryChan, i)
	}
}

func main() {
	Sim()
}
