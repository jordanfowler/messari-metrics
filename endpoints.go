package main

import (
	"context"
	"log"
	"os"

	metrics "github.com/jordanfowler/messari-metrics/gen/metrics"
	messari "github.com/jordanfowler/messari-metrics/messari"
)

// Metrics service example implementation.
type metricssrvc struct {
	logger *log.Logger
}

// NewMetrics returns the metrics service implementation.
func NewMetrics(logger *log.Logger) metrics.Service {
	return &metricssrvc{logger}
}

// Asset implements asset endpoint.
func (s *metricssrvc) Asset(ctx context.Context, p *metrics.AssetPayload) (res *metrics.AssetResult, err error) {
	res = &metrics.AssetResult{
		Metric: &metrics.AssetMetrics{},
	}
	client := messari.NewClient(os.Getenv("MESSARI_API_KEY"))

	assetMetrics, err := client.GetAssetMetrics(ctx, *p.Slug, map[string]interface{}{
		"fields": []string{
			"market_data",
			"marketcap",
		},
	})
	if err != nil {
		s.logger.Printf("Error [GetAssetMetrics]: %s", err)
		return
	}

	res.Metric.AssetSlug = p.Slug
	res.Metric.Price = &assetMetrics.MarketData.PriceUSD
	res.Metric.Mktcap = &assetMetrics.MarketCap.CurrentMarketCapUSD
	res.Metric.Chg24hr = &assetMetrics.MarketData.PercentChangeUSDLast24Hours
	res.Metric.Vlm24hr = &assetMetrics.MarketData.VolumeLast24Hours

	return
}

// Aggregate implements aggregate endpoint.
func (s *metricssrvc) Aggregate(ctx context.Context, p *metrics.AggregatePayload) (res *metrics.AggregateResult, err error) {
	res = &metrics.AggregateResult{
		Metrics: []*metrics.AssetMetrics{},
	}
	client := messari.NewClient(os.Getenv("MESSARI_API_KEY"))

	assets, err := client.GetAllAssets(ctx, map[string]interface{}{
		"fields": []string{
			"market_data",
		},
	})
	if err != nil {
		s.logger.Printf("Error [GetAllAssets]: %s", err)
		return
	}

	for _, m := range assets {
		var row = metrics.AssetMetrics{}

		row.AssetSlug = &m.Slug
		row.Price = &m.Metrics.MarketData.PriceUSD
		row.Mktcap = &m.Metrics.MarketCap.CurrentMarketCapUSD
		row.Chg24hr = &m.Metrics.MarketData.PercentChangeUSDLast24Hours
		row.Vlm24hr = &m.Metrics.MarketData.VolumeLast24Hours

		s.logger.Printf("slug: %s", *row.AssetSlug)

		res.Metrics = append(res.Metrics, &row)
	}

	return
}
