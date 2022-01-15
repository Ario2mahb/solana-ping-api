package main

import (
	"time"
)

func launchWorkers(clusters []Cluster, slackCluster []Cluster) {
	for _, c := range clusters {
		go getPingWorker(c)
	}

	time.Sleep(30 * time.Second)

	for _, c := range slackCluster {
		go slackReportWorker(c)
	}

}

func getPingWorker(c Cluster) {
	log.Info(">> Solana Ping Worker for ", c, " start!")
	for {
		startTime := time.Now().UTC().Unix()
		result := GetPing(c)
		endTime1 := time.Now().UTC().Unix()
		result.TakeTime = int(endTime1 - startTime)
		addRecord(result)
		endTime2 := time.Now().UTC().Unix()
		perPingTime := config.SolanaPing.PerPingTime
		waitTime := perPingTime - (endTime2 - startTime)
		if waitTime > 0 {
			time.Sleep(time.Duration(waitTime) * time.Second)
		}
	}
}

//GetPing  Do the solana ping and return ping result, return error is in PingResult.Error
func GetPing(c Cluster) PingResult {
	result := PingResult{Hostname: config.HostName, Cluster: string(c)}
	output, err := solanaPing(c)
	if err != nil {
		log.Error("GetPing ping error:", err)
		result.Error = err.Error()
		return result
	}
	err = result.parsePingOutput(output)
	if err != nil {
		log.Error("GetPing parse output error:", err)
		result.Error = err.Error()
		return result
	}
	return result
}

var lastReporUnixTime int64

func slackReportWorker(c Cluster) {
	log.Info(">> Slack Report Worker for ", c, " start!")
	for {
		if lastReporUnixTime == 0 {
			lastReporUnixTime = time.Now().UTC().Unix() - int64(config.Slack.ReportTime)
			log.Info("reconstruct lastReport time=", lastReporUnixTime, "time now=", time.Now().UTC().Unix())
		}
		data := getAfter(c, lastReporUnixTime)
		if len(data) <= 0 { // No Data
			log.Warn(c, " getAfter return empty")
			time.Sleep(30 * time.Second)
			continue
		}
		lastReporUnixTime = time.Now().UTC().Unix()
		stats := generateData(data)
		payload := SlackPayload{}
		payload.ToPayload(c, data, stats)
		err := SlackSend(config.Slack.WebHook, &payload)
		if err != nil {
			log.Error("SlackSend Error:", err)
		}

		time.Sleep(time.Duration(config.Slack.ReportTime) * time.Second)
	}
}