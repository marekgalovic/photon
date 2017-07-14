package metrics

import (
    "time";

    log "github.com/Sirupsen/logrus"
)

type StdoutClient struct {}

func (c *StdoutClient) Gauge(name string, value float64, tags []string, rate float64) error {
    log.WithFields(log.Fields{"name": name, "value": value, "tags": tags, "rate": rate}).Info("Metrics[gauge]")
    return nil
}

func (c *StdoutClient) Count(name string, value int64, tags []string, rate float64) error {
    log.WithFields(log.Fields{"name": name, "value": value, "tags": tags, "rate": rate}).Info("Metrics[count]")
    return nil
}

func (c *StdoutClient) Incr(name string, tags []string, rate float64) error {
    log.WithFields(log.Fields{"name": name, "tags": tags, "rate": rate}).Info("Metrics[incr]")
    return nil
}

func (c *StdoutClient) Histogram(name string, value float64, tags []string, rate float64) error {
    log.WithFields(log.Fields{"name": name, "value": value, "tags": tags, "rate": rate}).Info("Metrics[histogram]")
    return nil
}

func (c *StdoutClient) Timing(name string, value time.Duration, tags []string, rate float64) error {
    log.WithFields(log.Fields{"name": name, "value": value, "tags": tags, "rate": rate}).Info("Metrics[timing]")
    return nil
}
