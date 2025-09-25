// Using Go's built in JSON support for encoding and decoding
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//Encode and Decode JSON using customs types
type firstResponse struct {
	Page int
	Spices []string
}

//Encode and Decode JSON fields to be exported
type secondResponse struct {
	Page int `json:"page"`
	Spices []string `json:"spices"`
}


func main() {

	fmt.Println(strings.ToUpper("JSON"))
	//Encoding basic data types to JSON. 
	//Use Atomic values to encode to JSON

	//Boolean
	intoBoolean, _ := json.Marshal(true)
	fmt.Println(string(intoBoolean))

	//Integer
	integerBoolean, _ := json.Marshal(1)
	fmt.Println(string(integerBoolean))

	//Float
	floatBoolean, _ := json.Marshal(3.14)
	fmt.Println(floatBoolean)

	//String
	stringBoolean, _ := json.Marshal("Graphic")
	fmt.Println(string(stringBoolean))

	//Slices and Maps
	//Slice
	spiceSlice :=[]string{"rosemary", "cinnamon", "coriander", "basil", "clove"}
	sliceBoolean, _ := json.Marshal(spiceSlice)
	fmt.Println(string(sliceBoolean))

	//Map
	flowerMap := map[string]int{"Lilies": 2, "Roses": 3, "Lavender": 5, "Carnations": 7, "Orchids": 8}
	mapBoolean, _ := json.Marshal(flowerMap)
	fmt.Println(string(mapBoolean))

	//Automatically encode custom data types using JSON Package
	responseOne := &firstResponse {
		Page: 1,
		Spices: []string{"Cinnamon", "Basil", "Coriander", "Rosemary", "Clove"},
	}
	responseA, _ := json.Marshal(responseOne)
	fmt.Println(string(responseA))

	//Customise encoded JSON key names using tags on declared struct fields
	responseThree := &secondResponse {
		Page: 1,
		Spices: []string{"Cinnamon", "Basil", "Coriander", "Rosemary", "Clove"}}
	responseB, _ := json.Marshal(responseThree)
	fmt.Println(string(responseB))
	
	//Decode JSON data into Go values
	decodeBytes := []byte(`{"digits":7.15, "writings":["arya","blade"]}`)

	//Create variable to decode JSON data into. Maps are useful in storing arbitrary data types.
	var data map[string]interface{}

	//Decode the decodeBytes and data variables declared above.
	if err := json.Unmarshal(decodeBytes, &data); 
	err != nil {
		panic(err)
	}
	fmt.Println(data)

	//Convert values into relevant types so the values in the decoded map can be used
	numbers := data["digits"].(float64)
	fmt.Println(numbers)

	//Use a series of comversions to access nested data
	stringSlice := data["writings"].([]interface{})  // Changed from 'strings' to 'stringSlice'
	stringOne := stringSlice[0].(string)             // Updated reference
	fmt.Println(stringOne)

	//To add additional type-safety and limit type assertions, decode JSON into custom data types

	stringCollection := `{"page": 1, "spices" : ["Cinnamon", "Basil", "Coriander", "Rosemary", "Clove"]}`
	response := secondResponse{}
	json.Unmarshal([]byte(stringCollection), &response)
	fmt.Println(response)
	fmt.Println(response.Spices[0])

	//Stream JSON encodings to os Writers like os.Stdout or HTTP response body
	encodeStream := json.NewEncoder(os.Stdout)
	encodeMap := map[string]int{"Cinnamon": 2, "Basil": 3, "Coriander": 4, "Rosemary": 6, "Clove": 7}
	encodeStream.Encode(encodeMap)

	//Use json.Decoder to stream reads from os Readers like os.Stdin and HTTP Request bodies
	decodeStream := json.NewDecoder(strings.NewReader(stringCollection))

	finalResponse := secondResponse{}
	decodeStream.Decode(&finalResponse)
	fmt.Println(finalResponse)

}