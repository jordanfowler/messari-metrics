// Code generated by goa v3.2.6, DO NOT EDIT.
//
// metrics HTTP server types
//
// Command:
// $ goa gen github.com/jordanfowler/messari-metrics/design

package server

import (
	metrics "github.com/jordanfowler/messari-metrics/gen/metrics"
)

// AssetResponseBody is the type of the "metrics" service "asset" endpoint HTTP
// response body.
type AssetResponseBody AssetMetricsResponseBody

// AggregateResponseBody is the type of the "metrics" service "aggregate"
// endpoint HTTP response body.
type AggregateResponseBody AggregateMetricsResponseBody

// AssetMetricsResponseBody is used to define fields on response body types.
type AssetMetricsResponseBody struct {
	// Asset slug
	AssetSlug *string `form:"assetSlug,omitempty" json:"assetSlug,omitempty" xml:"assetSlug,omitempty"`
	// Current spot price in USD
	Price *float64 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
	// Volume traded over last 24 hours
	Vlm24hr *float64 `form:"vlm24hr,omitempty" json:"vlm24hr,omitempty" xml:"vlm24hr,omitempty"`
	// Change in price over last 24 hours
	Chg24hr *float64 `form:"chg24hr,omitempty" json:"chg24hr,omitempty" xml:"chg24hr,omitempty"`
	// Market cap of asset
	Mktcap *float64 `form:"mktcap,omitempty" json:"mktcap,omitempty" xml:"mktcap,omitempty"`
}

// AggregateMetricsResponseBody is used to define fields on response body types.
type AggregateMetricsResponseBody struct {
	// Aggregation name, e.g. tag, sector, etc.
	AggName *string `form:"aggName,omitempty" json:"aggName,omitempty" xml:"aggName,omitempty"`
	// Aggregation value, e.g. DeFi, FinTech, etc.
	AggValue *string `form:"aggValue,omitempty" json:"aggValue,omitempty" xml:"aggValue,omitempty"`
	// Current spot price in USD
	Price *float64 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
	// Volume traded over last 24 hours
	Vlm24hr *float64 `form:"vlm24hr,omitempty" json:"vlm24hr,omitempty" xml:"vlm24hr,omitempty"`
	// Change in price over last 24 hours
	Chg24hr *float64 `form:"chg24hr,omitempty" json:"chg24hr,omitempty" xml:"chg24hr,omitempty"`
	// Market cap of asset
	Mktcap *float64 `form:"mktcap,omitempty" json:"mktcap,omitempty" xml:"mktcap,omitempty"`
}

// NewAssetResponseBody builds the HTTP response body from the result of the
// "asset" endpoint of the "metrics" service.
func NewAssetResponseBody(res *metrics.AssetResult) *AssetResponseBody {
	body := &AssetResponseBody{
		AssetSlug: res.Metrics.AssetSlug,
		Price:     res.Metrics.Price,
		Vlm24hr:   res.Metrics.Vlm24hr,
		Chg24hr:   res.Metrics.Chg24hr,
		Mktcap:    res.Metrics.Mktcap,
	}
	return body
}

// NewAggregateResponseBody builds the HTTP response body from the result of
// the "aggregate" endpoint of the "metrics" service.
func NewAggregateResponseBody(res *metrics.AggregateResult) *AggregateResponseBody {
	body := &AggregateResponseBody{
		AggName:  res.Metrics.AggName,
		AggValue: res.Metrics.AggValue,
		Price:    res.Metrics.Price,
		Vlm24hr:  res.Metrics.Vlm24hr,
		Chg24hr:  res.Metrics.Chg24hr,
		Mktcap:   res.Metrics.Mktcap,
	}
	return body
}

// NewAssetPayload builds a metrics service asset endpoint payload.
func NewAssetPayload(slug string) *metrics.AssetPayload {
	v := &metrics.AssetPayload{}
	v.Slug = &slug

	return v
}

// NewAggregatePayload builds a metrics service aggregate endpoint payload.
func NewAggregatePayload(tag *string, sector *string) *metrics.AggregatePayload {
	v := &metrics.AggregatePayload{}
	v.Tag = tag
	v.Sector = sector

	return v
}
