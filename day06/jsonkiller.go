package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// tags are used to identify attributes in json
type Person struct {
	Name string `json:"name`
	Age  int    `json:"age`
}

func main() {

	someJson := `{"name" : "zahid","age":21}`

	pu := Person{}

	err := json.Unmarshal([]byte(someJson), &pu)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("pu= %+v\n", pu)

}
