package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// ExampleRequestData 是发送请求时使用的JSON数据结构
type ExampleRequestData struct {
	Username string `json:"username"`
	Password int    `json:"password"`
	// ... 其他字段
}

func TestConcurrentRequests(t *testing.T) {
	const (
		requestsCount = 100
		targetURI     = "http://127.0.0.1:3000/api/v1/users/add"
		requestType   = "POST" // 或 "GET"、"PUT" 等
	)

	// 示例请求数据

	for i := 0; i < requestsCount; i++ {
		requestBody := ExampleRequestData{
			Username: fmt.Sprintf("zhangsan%d", i),
			Password: 10000+i,
			// ... 初始化其他字段
		}

		// 将请求数据编码为JSON
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("Failed to marshal request body: %v", err)
		}
		req, err := http.NewRequest(requestType, targetURI, bytes.NewReader(jsonBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json") // 设置请求头 Content-Type 为 application/json

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("Failed to send request %d: %v", i, err)
			continue
		}
		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Failed to read response body for request %d: %v", i, err)
			continue
		}

		// 根据需要验证响应状态码、响应内容等
		// assert.Equal(t, http.StatusOK, resp.StatusCode, "Unexpected status code for request %d", i)
		// ... 其他验证逻辑

		// 或者，解析响应JSON并进行验证
		var responseData map[string]interface{}
		err = json.Unmarshal(respBody, &responseData)
		if err != nil {
			t.Errorf("Failed to unmarshal response body for request %d: %v", i, err)
			continue
		}
		// ... 验证 responseData

		// 输出进度（可选）
		t.Logf("Sent request %d/%d", i+1, requestsCount)
	}
}

func BenchmarkTestConcurrentRequests(t *testing.B) {
	t.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	
	const (
		requestsCount = 1000
		targetURI     = "http://127.0.0.1:3000/api/v1/users/add"
		requestType   = "POST" // 或 "GET"、"PUT" 等
	)

	// 示例请求数据
t.StartTimer() //重新开始时间
	for i := 0; i < requestsCount; i++ {
		requestBody := ExampleRequestData{
			Username: fmt.Sprintf("zhangsan%d", i),
			Password: 10000+i,
			// ... 初始化其他字段
		}

		// 将请求数据编码为JSON
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("Failed to marshal request body: %v", err)
		}
		req, err := http.NewRequest(requestType, targetURI, bytes.NewReader(jsonBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json") // 设置请求头 Content-Type 为 application/json

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("Failed to send request %d: %v", i, err)
			continue
		}
		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Failed to read response body for request %d: %v", i, err)
			continue
		}

		// 根据需要验证响应状态码、响应内容等
		// assert.Equal(t, http.StatusOK, resp.StatusCode, "Unexpected status code for request %d", i)
		// ... 其他验证逻辑

		// 或者，解析响应JSON并进行验证
		var responseData map[string]interface{}
		err = json.Unmarshal(respBody, &responseData)
		if err != nil {
			t.Errorf("Failed to unmarshal response body for request %d: %v", i, err)
			continue
		}
		// ... 验证 responseData

		// 输出进度（可选）
		t.Logf("Sent request %d/%d", i+1, requestsCount)
	}
}
