package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	CourseName string   `json:"courseName"`     // json:"coursename" is used to change the key name
	Price      int      `json:"-"`              // used to hide the key
	Author     string   `json:"author"`         // jused to change the key name
	Tags       []string `json:"tags,omitempty"` // omitempty is used to hide the key if it is empty
}

func main() {
	c := []Course{
		{
			CourseName: "Go",
			Price:      100,
			Author:     "Vinit",
			Tags:       []string{"Go", "Backend"},
		},

		{
			CourseName: "NodeJS",
			Price:      200,
			Author:     "Jay",
			Tags:       []string{"NodeJS", "Backend"},
		},
	}

	data := encodedJSON(c)
	fmt.Println("Normal:", c)
	fmt.Println("JSON:", string(data))

	decodedData := decodedJSON([]byte(`
	{
		"courseName": "Go",
		"price": 100,
		"author":"Vinit"
	}
	`))
	if decodedData == nil {
		return
	}
	fmt.Println("Decoded: ", decodedData)
}

func encodedJSON(data interface{}) []byte {
	// jsonData, err := json.Marshal(data)
	jsonData, err := json.MarshalIndent(data, "", "\t")
	fetal(err)
	return jsonData
}

func decodedJSON(data []byte) interface{} {
	var result Course
	isValid := json.Valid(data)
	if !isValid {
		fmt.Println("Invalid JSON")
		return nil
	}

	err := json.Unmarshal(data, &result)
	fetal(err)

	// in some cases we need to decode the data into a map
	var myMap map[string]interface{}
	json.Unmarshal(data, &myMap)
	fmt.Println("Mapped: ", myMap["courseName"])
	return result

}

func fetal(err error) {
	if err != nil {
		panic(err)
	}
}
