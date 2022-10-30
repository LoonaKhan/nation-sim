package name_gen

import (
	"fmt"
	"math/rand"
	"nsim/utils/json"
	"time"
)

// structs for our json data
type CountryName struct {
	Std string `json:"std"`
	Alt string `json:"alt"`
}

type CNameOps struct {
	Prefixes    []string      `json:"Prefixes"`
	Descriptors []string      `json:"Descriptors"`
	Suffixes    []string      `json:"Suffixes"`
	Names       []CountryName `json:"Names"`
}

type PeopleNames struct {
	FirstNames  []string `json:"firstNames"`
	MiddleNames []string `json:"middleNames"`
	LastNames   []string `json:"lastNames"`
}

var countryOps = json.Read[CNameOps]("nsim/names/countries.json")
var cNames = countryOps.Names
var cPrefixes = countryOps.Prefixes
var cSuffixes = countryOps.Suffixes
var cDescriptors = countryOps.Descriptors

// todo: categorize names based on ethnicity?
var pplNames = json.Read[PeopleNames]("nsim/names/people.json")
var firstNames = pplNames.FirstNames
var middleNames = pplNames.MiddleNames
var lastNames = pplNames.LastNames

func ChooseCountryName() string {
	/*
		chooses whether to use a prefix, a suffix or neither.
		all equally likely
		if a suffix is used, use the country's alt name
		otherwise use the standard name
		there is a small chance of adding "and" and another name
		also a small chance of adding a descriptor in any scenario
	*/
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	Name := ""

	switch r.Intn(3) {
	case 0: // use a prefix. ex: "The great repulic of frankia and angland"
		Name += "The "

		if r.Intn(10) == 1 { //adds a descriptor
			Name += cDescriptors[r.Intn(len(cDescriptors))] + " "
		}

		Name += fmt.Sprintf("%s %s", // "republic of frankia"
			cPrefixes[r.Intn(len(cPrefixes))],
			cNames[r.Intn(len(cNames))].Std)

		if r.Intn(10) == 1 { // add another name
			Name += fmt.Sprintf(" and %s", cNames[r.Intn(len(cNames))].Std) // "frankish and "
		}

		break

	case 1: // use a suffix. ex: "The great frankish and anglish republic"
		Name += "The "

		if r.Intn(10) == 1 { //adds a descriptor
			Name += cDescriptors[r.Intn(len(cDescriptors))] + " "
		}

		if r.Intn(10) == 1 { // add another name
			Name += fmt.Sprintf("%s and ", cNames[r.Intn(len(cNames))].Alt) // "frankish and "
		}

		Name += fmt.Sprintf("%s %s", // "anglish republic"
			cNames[r.Intn(len(cNames))].Alt,
			cSuffixes[r.Intn(len(cSuffixes))])

		break

	case 2: // use neither
		Name += cNames[r.Intn(len(cNames))].Std
		break
	}

	return Name
}

func ChoosePersonName() string {
	/*
		1 in 5 chance of adding a middle name
		very small chance of adding "of [cName]"
	*/ /*
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		name := firstNames[r.Intn(len(firstNames))] + " "

		if r.Intn(10) == 1 { // 1 in 10 chance of giving the pop a middle name
			name += middleNames[r.Intn(len(middleNames))] + " "
		}

		name += lastNames[r.Intn(len(lastNames))]*/

	name := "generic citizen name"

	return name
}
