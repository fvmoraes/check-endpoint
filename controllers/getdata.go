package controllers

import (
	"check-endpoint/logs"
	"io/ioutil"
	"strings"
	"time"
)

func GetInputArgs(file string, repeat, interval int) {
	//Args handling
	urlsList, err := ioutil.ReadFile(file)
	if err != nil {
		logs.LogApplicationError(err)
	}
	numberOfRepetitions := repeat
	timeBetweenRequests := interval
	//Loop to separate urls from passed list
	for numberOfRepetitions > 0 {
		urls := strings.Split(string(urlsList), "\n")
		for _, url := range urls {
			go GetHttpResponse(url, numberOfRepetitions)
		}
		numberOfRepetitions -= 1
		time.Sleep(time.Millisecond * time.Duration(timeBetweenRequests))
	}
}
