package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	test01()
	test02()
}

// 测试 /test
// curl -X POST -d '{"msg":"hello"}' -H "Content-Type:application/json" "http://127.0.0.1:10001/test" -v
func test01() {
	// 定义请求的 URL
	url := "http://127.0.0.1:10001/test"

	// 定义请求的 JSON 数据
	jsonData := []byte(`{"msg":"hello"}`)

	// 创建一个新的 POST 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应状态和响应体
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}

// 测试 /test/caculator
// curl -X POST -d '{"a":"1.0","b":"1.0"}' -H "Content-Type:application/json" "http://127.0.0.1:10001/test/caculator" -v
func test02() {
	// 定义请求的 URL
	url := "http://127.0.0.1:10001/test/caculator"

	// 定义请求的 JSON 数据
	jsonData := []byte(`{"a":"1", "b":"1", "op":"ADD"}`)

	// 创建一个新的 POST 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应状态和响应体
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}