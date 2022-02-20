package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
接口地址：
"http://localhost:8080/controller?appId=1"
*/
//get请求
func HttpGet() {
	resp, err := http.Get("http://localhost:8080/controller")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	var response Reponse
	if err := json.Unmarshal(body, &response); err == nil {
		fmt.Println(response)
		fmt.Println(response.Msg == "成功")
	} else {
		fmt.Println(err)
	}

}

/*返回参数结构*/
type GetIndexOverviewData struct {
}

/*返回结构*/
type Reponse struct {
	Code int                  `json:"code`
	Data GetIndexOverviewData `json:"data"`
	Msg  string               `json:"msg"`
}
