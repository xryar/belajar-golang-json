package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"p001", "name":"Macbook Pro", "price": 3000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapDecode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "p001",
		"name":  "Ipong 16 Pro Max",
		"price": 16000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
