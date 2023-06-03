package controllers

import (
	"check-endpoint/logs"
	"sync"
)

var (
	statisticsMap = make(map[string]Statistics)
	mutex         = &sync.Mutex{}
)

type Statistics struct {
	TotalResponseTime      int
	SuccessStatusCodeCount int
	TotalCount             int
}

func SaveStatistics(responseTime, StatusCode int, endpointUrl string) {
	mutex.Lock()
	defer mutex.Unlock()

	statistics, ok := statisticsMap[endpointUrl]
	if !ok {
		statistics = Statistics{}
	}

	if StatusCode >= 200 && StatusCode < 300 {
		statistics.SuccessStatusCodeCount++
	}
	statistics.TotalResponseTime += responseTime
	statistics.TotalCount++

	statisticsMap[endpointUrl] = statistics
}

func CalculateStatistics() {
	mutex.Lock()
	defer mutex.Unlock()

	for endpointUrl, statistics := range statisticsMap {
		averageResponseTime := statistics.TotalResponseTime / statistics.TotalCount
		successPercentage := float64(statistics.SuccessStatusCodeCount) / float64(statistics.TotalCount) * 100.0
		logs.LogApplicationStatistics(endpointUrl, averageResponseTime, successPercentage)
	}
}
