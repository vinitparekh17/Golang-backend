package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// doPostRequest("http://localhost:8080/post")
	doPostFormRequest("http://localhost:8080/postform")
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

func doPostFormRequest(link string) {
	data := url.Values{}
	data.Set("coursename", "Go")
	data.Set("price", "100")
	data.Set("author", "Vinit")

	res, err := http.PostForm(link, data)
	fetal(err)

	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Content:", string(append(content)))
}
