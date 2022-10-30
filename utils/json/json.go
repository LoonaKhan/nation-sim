package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Read[T any](path string) T { // uses a template to determine the structure of the json
	file, err := os.Open(path) // open the json file

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var data T
	byteValue, _ := ioutil.ReadAll(file)

	json.Unmarshal([]byte(byteValue), &data)

	return data

}
