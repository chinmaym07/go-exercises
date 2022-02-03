package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func PerformGetRequest() {
	const myUrl string = "http://localhost:3000/get"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder

	dataContent, err := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(dataContent)
	fmt.Println("ByteCount is", byteCount)
	fmt.Println("Response String", responseString.String())

	// fmt.Println("Data content", dataContent)
	// fmt.Println("Response Body", string(dataContent))
}

func PerformPostRequest() {
	const myUrl string = "http://localhost:3000/post"
	// fake json payload
	requestBody := strings.NewReader(`
		{
			"name": "Chinmay",
			"course": "golang",
			"price": 200,
			"platform": "coursera"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)
	dataContent, err := ioutil.ReadAll(response.Body)
	var responseContent strings.Builder
	data, _ := responseContent.Write(dataContent)
	fmt.Println(data)
	fmt.Println("Response Content", responseContent.String())

	fmt.Println("Data Content", string(dataContent))

}
func PerformPostFormEncodedRequest() {
	const myUrl string = "http://localhost:3000/postform"
	// fake form data
	requestBody := &url.Values{}
	requestBody.Add("firstname", "Chinmay")
	requestBody.Add("lastname", "Mehta")
	requestBody.Add("course", "golang")
	requestBody.Add("price", "200")
	requestBody.Add("platform", "coursera")
	requestBody.Add("email", "test@gmail.com")

	response, err := http.PostForm(myUrl, *requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)
	dataContent, err := ioutil.ReadAll(response.Body)
	var responseContent strings.Builder
	data, _ := responseContent.Write(dataContent)
	fmt.Println(data)
	fmt.Println("Response Content", responseContent.String())

	fmt.Println("Data Content", string(dataContent))

}

func main() {
	fmt.Println("Server Request in Golang")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormEncodedRequest()
}
