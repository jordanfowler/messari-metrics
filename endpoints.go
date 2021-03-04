package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	res = &metrics.AssetResult{}
	client := messari.NewClient(os.Getenv("MESSARI_API_KEY"))
	// res.Metric, err = fetchAssetMetrics(ctx, client, *p.Slug)
	amRes, err := client.GetAssetMetrics(ctx, *p.Slug, map[string]interface{}{
		"fields": []string{
			"market_data",
			"marketcap",
		},
	})
	if err != nil {
		return
	}

	res.Metric = &metrics.AssetMetrics{
		AssetSlug: p.Slug,
		Price:     &amRes.MarketData.PriceUSD,
		Mktcap:    &amRes.MarketCap.CurrentMarketCapUSD,
		Chg24hr:   &amRes.MarketData.PercentChangeUSDLast24Hours,
		Vlm24hr:   &amRes.MarketData.VolumeLast24Hours,
	}

	return
}

// Aggregate implements aggregate endpoint.
func (s *metricssrvc) Aggregate(ctx context.Context, p *metrics.AggregatePayload) (res *metrics.AggregateResult, err error) {
	res = &metrics.AggregateResult{
		Metrics: []*metrics.AssetMetrics{},
	}

	var assetSlugs []string
	if p.Sector != nil {
		assetSlugs = GetAssetCacheSlugs("sectors", *p.Sector)
		s.logger.Println(fmt.Sprintf("Aggregate payload sector=%s slugs=%v", *p.Sector, assetSlugs))
	} else if p.Tag != nil {
		assetSlugs = GetAssetCacheSlugs("tags", *p.Tag)
		s.logger.Println(fmt.Sprintf("Aggregate payload tag=%s slugs=%v", *p.Tag, assetSlugs))
	}

	var aggregateCh = make(chan *metrics.AssetResult, len(assetSlugs))
	var errCh = make(chan error)

	if len(assetSlugs) > 0 {
		for _, slug := range assetSlugs {
			go func(slug string) {
				ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
				defer cancel()
				res, err := s.Asset(ctx, &metrics.AssetPayload{Slug: &slug})
				if err != nil {
					s.logger.Println("fetchAssetMetrics error:", err)
					errCh <- err
					return
				}
				s.logger.Println("fetchAssetMetrics:", res)
				aggregateCh <- res
			}(slug)
		}

		for {
			select {
			case ar := <-aggregateCh:
				s.logger.Println("aggregateCh ar:", ar)
				res.Metrics = append(res.Metrics, ar.Metric)
			case err = <-errCh:
				s.logger.Println("aggregateCh error:", err)
				return
			default:
				if len(res.Metrics) == len(assetSlugs) {
					return
				}
			}
		}
	}

	return
}
