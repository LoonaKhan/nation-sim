package main

import (
	"fmt"
	"nsim/nsim"
)

func GenCountryNameTest() {
	fmt.Println(nsim.ChooseCountryName())
}

func main() {
	for i := 0; i < 10; i++ {
		GenCountryNameTest()
	}
}
