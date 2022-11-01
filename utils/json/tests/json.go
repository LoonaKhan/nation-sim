package main

import (
	"fmt"
	"nsim/nsim/globvars"
	"nsim/nsim/names"
	"nsim/nsim/ppl"
	"nsim/utils/json"
)

func ReadJsonTest(path string) names.CNameOps { // we test it with a country name
	return json.Read[names.CNameOps](path)
}

func ReadJobsTest(path string) map[string]ppl.Job {
	return json.Read[map[string]ppl.Job](path)
}

func ReadGlobvarsTest(path string) globvars.Globvars {
	return json.Read[globvars.Globvars](path)
}

func main() {
	//fmt.Println(ReadJsonTest("nsim/names/countries.json"))
	//fmt.Println(ReadJobsTest("nsim/pop/jobs.json"))
	fmt.Println(ReadGlobvarsTest("nsim/globvars/globvars.json"))
}
