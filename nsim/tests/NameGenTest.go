package main

import (
	"fmt"
	"nsim/nsim"
)

func genCountryNameTest() {
	fmt.Println(nsim.ChooseCountryName())
}

func genPeopleNameTest() {
	fmt.Println(nsim.ChoosePersonName())
}

func main() {
	for i := 0; i < 10; i++ {
		//genCountryNameTest()
		genPeopleNameTest()
	}
}
