package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "db_backup",
		Name:      "last_completion_timestamp_seconds",
		Help:      "The timestamp of the last successful completion of a DB backup.",
	})

	jobCount := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "db_backup",
		Name:      "job_finish_count",
		Help:      "The total jobs finish",
	})

	histTimeCost := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "db_backup",
		Name:      "histogram_timecost_seconds",
		Help:      "The time cost every time to db backup",
		Buckets: []float64{
			4, 8, 16, 32, 64,
		},
	})

	summaryTimeCost := prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: "db_backup",
		Name:      "summary_timecost_seconds",
		Help:      "The time cost every time to db backup",
		Objectives: map[float64]float64{
			0.95: 0,
			0.8:  0,
			0.5:  0,
		},
	})

	// update metrics
	for i := 1; i < 20; i++ {
		jobCount.Add(1)
		histTimeCost.Observe(float64(i) * 2)
		summaryTimeCost.Observe(float64(i) * 2)
	}
	completionTime.SetToCurrentTime()

	// update username/password to real telmetry tower account
	pusher := push.New("https://io.telemetrytower.com/pushgateway", "db_backup").
		Grouping("instance", "cluster01").
		BasicAuth("username", "password")

	// register metrics
	pusher.Collector(completionTime)
	pusher.Collector(jobCount)
	pusher.Collector(histTimeCost)
	pusher.Collector(summaryTimeCost)

	// push data to telemetry tower
	if err := pusher.Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}
