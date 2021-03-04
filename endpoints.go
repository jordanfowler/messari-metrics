package main

import (
	"context"
	"log"
	"os"

	metrics "github.com/jordanfowler/messari-metrics/gen/metrics"
	messari "github.com/jordanfowler/messari-metrics/messari"
)

// metrics service example implementation.
// The example methods log the requests and return zero values.
type metricssrvc struct {
	logger *log.Logger
}

// NewMetrics returns the metrics service implementation.
func NewMetrics(logger *log.Logger) metrics.Service {
	return &metricssrvc{logger}
}

// Asset implements asset endpoint.
func (s *metricssrvc) Asset(ctx context.Context, p *metrics.AssetPayload) (res *metrics.AssetResult, err error) {
	res = &metrics.AssetResult{}
	client := messari.NewClient(os.Getenv("MESSARI_API_KEY"))

	metrics, err := client.GetAssetMetrics(ctx, *p.Slug, map[string]interface{}{
		"fields": []string{
			"market_data",
			"marketcap",
		},
	})
	if err != nil {
		s.logger.Printf("Error [GetAssetMetrics]: %s", err)
		return
	}

	res.Price = &metrics.MarketData.PriceUSD
	res.Mktcap = &metrics.MarketCap.CurrentMarketCapUSD
	res.Chg24hr = &metrics.MarketData.PercentChangeUSDLast24Hours
	res.Vlm24hr = &metrics.MarketData.VolumeLast24Hours

	return
}

// Aggregate implements aggregate endpoint.
func (s *metricssrvc) Aggregate(ctx context.Context, p *metrics.AggregatePayload) (res *metrics.AggregateResult, err error) {
	res = &metrics.AggregateResult{}
	s.logger.Print("metrics.aggregate")
	return
}
