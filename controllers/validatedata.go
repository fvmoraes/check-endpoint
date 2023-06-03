package controllers

import (
	"check-endpoint/logs"
	"io"
	"net/http"
	"strconv"
	"time"
)

func GetHttpResponse(url string, numberOfRepetitions int) {
	//Starting the request time count and performing a GET on the received url
	startTime := time.Now()
	httpResponse, err := http.Get(url)
	if err != nil {
		logs.LogApplicationError(err)
	}
	FilterInformationFromHttpResponse(httpResponse, numberOfRepetitions, startTime)
}

func FilterInformationFromHttpResponse(httpResponse *http.Response, numberOfRepetitions int, startTime time.Time) {
	if httpResponseBody, err := io.ReadAll(httpResponse.Body); err != nil {
		logs.LogApplicationError(err)
	} else {
		endpointUrl := httpResponse.Request.URL.String()
		statusCode := httpResponse.Status
		responseTime := strconv.Itoa(int(time.Since(startTime).Milliseconds()))
		responseBody := string(httpResponseBody)
		SaveStatistics(int(time.Since(startTime).Milliseconds()), httpResponse.StatusCode, endpointUrl)
		logs.LogApplicationStdout(endpointUrl, statusCode, responseTime, numberOfRepetitions)
		logs.LogApplicationResponse(endpointUrl, statusCode, responseTime, responseBody)
	}
	defer httpResponse.Body.Close()
}
