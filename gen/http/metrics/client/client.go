// Code generated by goa v3.2.6, DO NOT EDIT.
//
// metrics client HTTP transport
//
// Command:
// $ goa gen github.com/jordanfowler/messari-metrics/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the metrics service endpoint HTTP clients.
type Client struct {
	// Asset Doer is the HTTP client used to make requests to the asset endpoint.
	AssetDoer goahttp.Doer

	// Aggregate Doer is the HTTP client used to make requests to the aggregate
	// endpoint.
	AggregateDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the metrics service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AssetDoer:           doer,
		AggregateDoer:       doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Asset returns an endpoint that makes HTTP requests to the metrics service
// asset server.
func (c *Client) Asset() goa.Endpoint {
	var (
		decodeResponse = DecodeAssetResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAssetRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AssetDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("metrics", "asset", err)
		}
		return decodeResponse(resp)
	}
}

// Aggregate returns an endpoint that makes HTTP requests to the metrics
// service aggregate server.
func (c *Client) Aggregate() goa.Endpoint {
	var (
		encodeRequest  = EncodeAggregateRequest(c.encoder)
		decodeResponse = DecodeAggregateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAggregateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AggregateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("metrics", "aggregate", err)
		}
		return decodeResponse(resp)
	}
}
