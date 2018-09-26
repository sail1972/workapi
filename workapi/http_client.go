package workapi

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// HTTPRequest http请求客户端
func HTTPRequest(url string, method string, body io.Reader) ([]byte, error) {

	client := &http.Client{}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err

	}

	if method == "POST" {
		request.Header.Set("Accept", "application/json")
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}

	//处理返回结果
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	buf, err := ioutil.ReadAll(response.Body)

	if response.StatusCode == http.StatusOK {
		return buf, nil
	}

	return nil, errors.New(fmt.Sprint("StatusCode=", response.StatusCode, " msg=", string(buf)))

}

// PostRequest POST 方式调用 HTTP 请求
func PostRequest(url string, body io.Reader) ([]byte, error) {
	return HTTPRequest(url, "POST", body)
}

// GetRequest GET 方式调用 HTTP 请求
func GetRequest(url string) ([]byte, error) {
	return HTTPRequest(url, "GET", nil)
}
