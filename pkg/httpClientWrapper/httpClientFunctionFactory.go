package httpClientWrapper

import (
	"net/http"
)

func BuildHttpClientFunction(url string, httpMethod string, bearer string) func() int {
	return func() int {
		request := createHttpRequest(url, httpMethod, bearer)

		client := &http.Client{}
		response, err := client.Do(request)
		response.Body.Close()
		if err != nil {
			return -1
		}

		return response.StatusCode
	}
}
