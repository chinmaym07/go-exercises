package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json in golang")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"ReactJs", 299, "coursera", "abcd@123", []string{"Web-development", "frontend"}},
		{"NodeJs", 399, "coursera", "abcd@123", []string{"Web-development", "backend"}},
		{"Full Stack", 399, "coursera", "abcd@123", nil},
	}
	// package this data as json data
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Final Json %s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`{
		"coursename": "NodeJs",
		"price": 399,
		"platform": "coursera",
		"tags": [
				"Web-development",
				"backend"
			]
		}`)
	var course1 course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Valid Json !!")
		err := json.Unmarshal(jsonDataFromWeb, &course1)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", course1)
	} else {
		fmt.Println("Json Invalid")
	}
	// Some cases you just want to add data to key value pair

	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("key is %v & value is %v & type of %T\n", k, v, v)
	}
}
