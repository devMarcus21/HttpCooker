package httpClientWrapper

import (
	"net/http"
)

func createHttpRequest(url string, httpMethod string, bearer string) *http.Request {
	request, _ := http.NewRequest(httpMethod, url, nil)

	if len(bearer) > 0 {
		request.Header.Add("Authorization", "Bearer "+bearer)
	}

	return request
}
