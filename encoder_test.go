package belajargolangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := &Customer{
		FirstName:  "Arya",
		MiddleName: "Rizki",
		LastName:   "Andaru",
	}

	encoder.Encode(customer)
}
