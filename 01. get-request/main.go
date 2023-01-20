package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	doGetRequest("http://localhost:8080/get")
}

func doGetRequest(url string) {
	res, err := http.Get(url)
	fetal(err)
	defer res.Body.Close()

	fmt.Println("StatusCode:", res.StatusCode)

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Content:", string(content))
}

func fetal(err error) {
	if err != nil {
		panic(err)
	}
}
