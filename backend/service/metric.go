package service

import (
	"context"
	"errors"
	"strconv"
)

type MetricStore interface {
	Inc(k string) error
	Last(string, int) (int, error)
}

type Metric interface {
	LastMetrics(ctx context.Context, n string, s string) (int, error)
}

type metric struct {
	ms MetricStore
}

func (m *metric) LastMetrics(ctx context.Context, metricName string, interval string) (int, error) {
	intervalNum, err := parseToUnix(interval)
	if err != nil {
		return 0, err
	}
	// TODO: dont ignore last metric error
	result, _ := m.ms.Last(metricName, intervalNum)
	return result, nil
}

func NewMetric(ms MetricStore) Metric {
	return &metric{
		ms: ms,
	}
}

func parseToUnix(s string) (int, error) {
	index := len(s) - 1
	numStr := s[:index]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, err
	}
	t := s[index:]
	if v, ok := metricType[t]; ok {
		return v * num, nil
	}
	return 0, errors.New("invalid metric type")
}

const (
	second = 1000
	minute = 60 * second
	hour   = 60 * minute
)

var metricType = map[string]int{
	"h": hour,
	"m": minute,
}
