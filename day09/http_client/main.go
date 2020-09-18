package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

// net/http Client

func main() {
	// resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=张三&age=18")
	// if err != nil {
	// 	fmt.Println("get url failed, err:", err)
	// 	return
	// }
	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "张三")
	data.Set("age", "22")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Println("get url failed, err:", err)
	// 	return
	// }

	// 请求不是特别频繁，用完就关闭
	// client := http.Client{
	// 	Transport: &http.Transport{
	// 		DisableKeepAlives: true,
	// 	},
	// }

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get url failed, err:", err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.Body failed, err:", err)
		return
	}
	fmt.Println(string(b))
}
