package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Json encoding tutorial")
	// EncodeJson()
	DecodeJson()
}

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReacrJS Bootcamp", 299, "LearnCodeOnline", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 199, "LearnCodeOnline", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline", "azi123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReacrJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// just in key value format

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v, and Type is: %T\n", k, v, v)
	}
}
