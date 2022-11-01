package redis

import (
	"backend/config"
	"backend/service"
	"errors"
	"time"

	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
)

type redis struct {
	retention int64
	c         *redistimeseries.Client
}

func (r *redis) Last(name string, interval int) (int, error) {
	ops := redistimeseries.NewRangeOptions().SetAggregation(redistimeseries.SumAggregation, interval)
	now := time.Now()
	delta := time.Millisecond * time.Duration(interval)
	from := now.Add(-delta)
	dp, err := r.c.RangeWithOptions(name, from.UnixMilli(), now.UnixMilli(), *ops)
	if err != nil {
		return 0, err
	}
	if len(dp) > 0 {
		return int(dp[0].Value), nil
	}
	return 0, errors.New("invalid query")
}

func (r *redis) Inc(k string) error {
	_, err := r.c.Info(k)
	if err != nil {
		r.c.CreateKeyWithOptions(k, redistimeseries.CreateOptions{
			Uncompressed:    false,
			RetentionMSecs:  0,
			Labels:          map[string]string{},
			ChunkSize:       0,
			DuplicatePolicy: "sum",
		})
	}
	// TODO: create before increment?
	_, err = r.c.AddAutoTs(k, 1)
	if err != nil {
		return err
	}
	return nil
}

func New(cfg config.Redis) service.MetricStore {

	client := redistimeseries.NewClient(cfg.URI, cfg.Name, cfg.Pass)

	return &redis{
		// delta:     time.Millisecond * time.Duration(cfg.Retention),
		c:         client,
		retention: cfg.Retention,
	}
}
