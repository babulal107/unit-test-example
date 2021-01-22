package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func printIfPerson(object interface{}) Person {
	person, ok := object.(Person) //type casting to expected struct
	if ok {
		fmt.Printf("Hello %s!\n", person.FirstName)
	}
	return person
}

func main() {
	person := Person{FirstName: "Babulal", LastName: "Choudhary"}
	resp := printIfPerson(person)
	fmt.Println("Person Struct 1 : ", resp)

	personObj := Person{}
	val := "{\"first_name\":\"Ankit\",\"last_name\":\"Kumbhar\"}"
	ba := []byte(val)
	_ = json.Unmarshal(ba, &personObj)

	resp2 := printIfPerson(personObj)
	fmt.Println("Person Struct 2 : ", resp2)
}
