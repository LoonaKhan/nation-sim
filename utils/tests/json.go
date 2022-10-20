package main

import (
	"fmt"
	"nsim/nsim"
	"nsim/utils"
)

func ReadJsonTest(path string) nsim.CNameOps { // we test it with a country name
	return utils.Read[nsim.CNameOps](path)

}

func main() {
	fmt.Println(ReadJsonTest("nsim/names/countries.json"))
}
