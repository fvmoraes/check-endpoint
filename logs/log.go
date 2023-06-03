package logs

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const myTimeFormat = "02/01/2006 15:04:05"

var logFile *os.File

func init() {
	var err error
	logFile, err = os.OpenFile("./check_endpoint.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
}

func LogApplicationStartInformation(command, file string, repeat, interval int) {
	logString := fmt.Sprintf("\n"+
		"--- Start Verification ---\n"+
		"--------------------------\n"+
		"COMMAND: %s\n"+
		"LIST WITH URLS: %s\n"+
		"REPETITIONS DEFINED: %d\n"+
		"TIME BETWEEN REQUESTS: %d\n"+
		"---",
		command, file, repeat, interval)
	fmt.Print(logString)
	WriteToLogFile(logString)
}

func LogApplicationStdout(endpointUrl, statusCode, responseTime string, numberOfRepetitions int) {
	isUrl := strings.Contains(endpointUrl, "http")
	if isUrl {
		logString := fmt.Sprintf("---\n"+
			"ENDPOINT URL: %s\n"+
			"STATUS CODE: %s\n"+
			"RESPONSE TIME: %s ms\n"+
			"REPEATS REMAINING: %d\n"+
			"---",
			endpointUrl, statusCode, responseTime, numberOfRepetitions)
		fmt.Print(logString)
	} else {
		logString := fmt.Sprintf("ERROR: URL (%s) not found", endpointUrl)
		fmt.Print(logString)
	}
}

func LogApplicationStatistics(endpointUrl string, averageResponseTime int, successPercentage float64) {
	logString := fmt.Sprintf("---\n"+
		"This is the result of your endpoint checks:\n"+
		"=> Url = %s\n"+
		"=> Average Response Time = %d ms\n"+
		"=> Success Percentage = %.2f%% \n"+
		"---",
		endpointUrl, averageResponseTime, successPercentage)
	fmt.Print(logString)
	WriteToLogFile(logString)
}

func LogApplicationResponse(endpointUrl, statusCode, responseTime, responseBody string) {
	logString := fmt.Sprintf("---\n"+
		"ENDPOINT URL: %s\n"+
		"STATUS CODE: %s\n"+
		"RESPONSE TIME: %s ms\n"+
		"CONTENT BODY: %s\n"+
		"---",
		endpointUrl, statusCode, responseTime, responseBody)
	WriteToLogFile(logString)
}

func LogApplicationError(logError error) {
	logString := fmt.Sprintf("---\nERROR: %s", logError.Error())
	WriteToLogFile(logString)
	fmt.Print(logString)
}

func WriteToLogFile(information string) {
	logEntry := time.Now().Format(myTimeFormat) + information + "\n"
	if _, err := logFile.WriteString(logEntry); err != nil {
		log.Println("Failed to write to log file:", err)
	}
}
