package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	/*resp, err := http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Printf("Get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read from resp.Body failed, err:%v\n", err)
		return
	}

	content := string(body)
	err = ioutil.WriteFile("liwenzhou.html", []byte(content), 0666)
	if err != nil {
		fmt.Println("Write failed, err:", err)
		return
	}*/

	apiUrl := "http://127.0.0.1:9090/get"
	// URL param
	data := url.Values{}
	data.Set("name", "Golang")
	data.Set("age", "13")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("get failed.err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed.err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
