package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := &Customer{
		FirstName:  "Arya",
		MiddleName: "Rizki",
		LastName:   "Andaru",
		Gender:     "Male",
		Age:        22,
		Married:    false,
		Hobbies:    []string{"Reading", "Gaming", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Arya","MiddleName":"Rizki","LastName":"Andaru","Age":22,"Gender":"Male","Married":false,"Hobbies":["Reading","Gaming","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := &Customer{
		FirstName: "Arya",
		Addresses: []Address{
			{
				Street:     "Jalan Acumalaka",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
			{
				Street:     "Jalan Malaka",
				Country:    "Indonesia",
				PostalCode: "9889",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Arya","MiddleName":"","LastName":"","Age":0,"Gender":"","Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Acumalaka","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Malaka","Country":"Indonesia","PostalCode":"9889"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Acumalaka","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Malaka","Country":"Indonesia","PostalCode":"9889"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplexEncode(t *testing.T) {
	addresses := &[]Address{
		{
			Street:     "Jalan Acumalaka",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "Jalan Malaka",
			Country:    "Indonesia",
			PostalCode: "9889",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
