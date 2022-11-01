package main

import (
	"fmt"
	"nsim/nsim/names"
)

func genCountryNameTest() {
	fmt.Println(names.ChooseCountryName())
}

func genPeopleNameTest() {
	fmt.Println(names.ChoosePersonName())
}

func main() {
	for i := 0; i < 10; i++ {
		//genCountryNameTest()
		genPeopleNameTest()
	}
}
