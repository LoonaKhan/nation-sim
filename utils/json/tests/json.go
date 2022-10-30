package main

import (
	"fmt"
	"nsim/nsim/name-gen"
	"nsim/nsim/pop"
	"nsim/utils/json"
)

func ReadJsonTest(path string) name_gen.CNameOps { // we test it with a country name
	return json.Read[name_gen.CNameOps](path)
}

func ReadJobsTest(path string) map[string]pop.Job {
	return json.Read[map[string]pop.Job](path)
}

func main() {
	//fmt.Println(ReadJsonTest("nsim/names/countries.json"))
	fmt.Println(ReadJobsTest("nsim/pop/jobs.json"))
}
