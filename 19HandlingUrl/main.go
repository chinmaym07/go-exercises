package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gdsfdvcdsfdse"

func main() {
	fmt.Println("Handling Urls in Golang")
	fmt.Println("Url, ", myUrl)

	// parsing the Url
	result, _ := url.Parse(myUrl)
	fmt.Println("Result Scheme", result.Scheme)
	fmt.Println("Result Host", result.Host)
	fmt.Println("Result Path", result.Path)
	fmt.Println("Result Port", result.Port())
	fmt.Println("Result Query", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Query Params %T\n", qparams)
	fmt.Println("Query Param\n", qparams)

	fmt.Println("Course Name", qparams["coursename"])
	fmt.Println("Payment ID", qparams["paymentid"])

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
