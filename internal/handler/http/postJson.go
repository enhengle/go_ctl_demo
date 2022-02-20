package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//params的post请求
func HttpParams() {
	resp, err := http.Post("http://localhost:8080/controller",
		"application/x-www-form-urlencoded", strings.NewReader("id=1,date=2021-11-11"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

}

//json的post请求
func HttpJson() {

	requestBody := fmt.Sprintln(`{
	}`)

	resp, err := http.Post("http://localhost:8080/controller",
		"application/json", strings.NewReader(requestBody))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

}

//form的post请求
func HttpFormDate() {

	requestBody := url.Values{}
	requestBody.Set("", "")
	formDataBody := []byte(requestBody.Encode())

	resp, err := http.Post("http://localhost:8080/controller",
		"application/json", bytes.NewReader(formDataBody))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

}
