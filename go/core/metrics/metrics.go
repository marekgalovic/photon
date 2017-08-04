package metrics

import (
    "fmt";
    "time";

    log "github.com/Sirupsen/logrus"
)

var metricsClient ClientInterface = &StdoutClient{}
var metricsBasename string = "serving"

type ClientInterface interface {
    Gauge(name string, value float64, tags []string, rate float64) error
    Count(name string, value int64, tags []string, rate float64) error
    Incr(name string, tags []string, rate float64) error
    Histogram(name string, value float64, tags []string, rate float64) error
    Timing(name string, value time.Duration, tags []string, rate float64) error
}

func SetClient(client ClientInterface) {
    metricsClient = client
}

func GetClient() ClientInterface {
    return metricsClient
}

func Gauge(name string, value float64, tags []string, rate float64) {
    if err := metricsClient.Gauge(fmt.Sprintf("%s.%s", metricsBasename, name), value, tags, rate); err != nil {
        log.Errorf("Failed to send gauge metric %s - %v", name, err)
    }
}

func Count(name string, value int64, tags []string, rate float64) {
    if err := metricsClient.Count(fmt.Sprintf("%s.%s", metricsBasename, name), value, tags, rate); err != nil {
        log.Errorf("Failed to send count metric %s - %v", name, err)
    }
}

func Incr(name string, tags []string, rate float64) {
    if err := metricsClient.Incr(fmt.Sprintf("%s.%s", metricsBasename, name), tags, rate); err != nil {
        log.Errorf("Failed to send incr metric %s - %v", name, err)
    }
}

func Histogram(name string, value float64, tags []string, rate float64) {
    if err := metricsClient.Histogram(fmt.Sprintf("%s.%s", metricsBasename, name), value, tags, rate); err != nil {
        log.Errorf("Failed to send histogram metric %s - %v", name, err)
    }
}

func Timing(name string, value time.Duration, tags []string, rate float64) {
    if err := metricsClient.Timing(fmt.Sprintf("%s.%s", metricsBasename, name), value, tags, rate); err != nil {
        log.Errorf("Failed to send timing metric %s - %v", name, err)
    }
}

func Runtime(name string, tags []string) func() {
    startAt := time.Now()
    
    return func() {
        Timing(name, time.Since(startAt), tags, 1)
    }
}
