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
type AssetResponseBody struct {
	// Current spot price in USD
	Price *float64 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
	// Volume traded over last 24 hours
	Vlm24hr *float64 `form:"vlm24hr,omitempty" json:"vlm24hr,omitempty" xml:"vlm24hr,omitempty"`
	// Change in price over last 24 hours
	Chg24hr *float64 `form:"chg24hr,omitempty" json:"chg24hr,omitempty" xml:"chg24hr,omitempty"`
	// Market cap of asset
	Mktcap *float64 `form:"mktcap,omitempty" json:"mktcap,omitempty" xml:"mktcap,omitempty"`
}

// AggregateResponseBody is the type of the "metrics" service "aggregate"
// endpoint HTTP response body.
type AggregateResponseBody struct {
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
		Price:   res.Price,
		Vlm24hr: res.Vlm24hr,
		Chg24hr: res.Chg24hr,
		Mktcap:  res.Mktcap,
	}
	return body
}

// NewAggregateResponseBody builds the HTTP response body from the result of
// the "aggregate" endpoint of the "metrics" service.
func NewAggregateResponseBody(res *metrics.AggregateResult) *AggregateResponseBody {
	body := &AggregateResponseBody{
		Price:   res.Price,
		Vlm24hr: res.Vlm24hr,
		Chg24hr: res.Chg24hr,
		Mktcap:  res.Mktcap,
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
func NewAggregatePayload(tags *string, sector *string) *metrics.AggregatePayload {
	v := &metrics.AggregatePayload{}
	v.Tags = tags
	v.Sector = sector

	return v
}
