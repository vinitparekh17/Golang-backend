package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	doPostRequest("http://localhost:8080/post")
}

func doPostRequest(url string) {
	reqData := strings.NewReader(`
	{
		"coursename": "Go",
		"price": 100,
		"author": "Vinit"
	}`)

	res, err := http.Post(url, "application/json", reqData)
	fetal(err)

	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Content:", string(content))
}

func fetal(err error) {
	if err != nil {
		panic(err)
	}
}
