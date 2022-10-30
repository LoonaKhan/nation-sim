package main

import (
	"fmt"
	"nsim/nsim/name-gen"
)

func genCountryNameTest() {
	fmt.Println(name_gen.ChooseCountryName())
}

func genPeopleNameTest() {
	fmt.Println(name_gen.ChoosePersonName())
}

func main() {
	for i := 0; i < 10; i++ {
		//genCountryNameTest()
		genPeopleNameTest()
	}
}
