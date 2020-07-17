package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post(params map[string]string)  {
	client := &http.Client{}
	data := make(map[string]interface{})
	data["name"] = "zhaofan"
	data["age"] = "23"
	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST","http://www.xxx.com",bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	defer resp.Body.close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func Get()  {
	var params map[string]string
	body := httpRequest("https://www.jianshu.com/p/4fbf529926ca", "GET", params, params)
	fmt.Println(string(body))
}

func httpRequest(url string, method string, params map[string]string, headers map[string]string) []byte {
	client := &http.Client{}
	var bytesData []byte
	var req *http.Request

	if method == "POST" {
		jsonData, _ := json.Marshal(params)
		bytesData = jsonData
		reqHttp,_ := http.NewRequest(method,url, bytes.NewReader(bytesData))
		req = reqHttp
	}

	if len(headers) > 0 {
		for key, value := range headers {
			req.Header.Add(key,value)
		}
	}

	resp,_ := client.Do(req)
	defer resp.Body.close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}