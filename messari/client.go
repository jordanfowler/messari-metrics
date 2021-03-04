package messari

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	// BaseURLV1 is v1 URL of Messare API
	BaseURLV1 = "https://data.messari.io/api/v1"
	// BaseURLV2 is v2 URL of Messare API
	BaseURLV2 = "https://data.messari.io/api/v2"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// AssetMetrics provide mapped access to asset metric response data
type AssetMetrics struct {
	MarketData struct {
		PriceUSD                    float64 `json:"price_usd"`
		VolumeLast24Hours           float64 `json:"volume_last_24_hours"`
		PercentChangeUSDLast24Hours float64 `json:"percent_change_usd_last_24_hours"`
	} `json:"market_data"`
	MarketCap struct {
		CurrentMarketCapUSD float64 `json:"current_marketcap_usd"`
	} `json:"marketcap"`
	MiscData struct {
		Tags   []string `json:"tags"`
		Sector []string `json:"sector"`
	} `json:"misc_data"`
}

// Asset provide mapped access to asset response data
type Asset struct {
	Slug    string       `json:"slug"`
	Metrics AssetMetrics `json:"metrics"`
}

// Client wraps calls to Messari API
type Client struct {
	apiKey     string
	HTTPClient *http.Client
}

// NewClient returns a new Client with a specified apiKey
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) fetch(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	fullResponse := successResponse{
		Data: v,
	}
	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}

	return nil
}

// GetAssetMetrics calls /api/v1/assets/{slug}/metrics
func (c *Client) GetAssetMetrics(ctx context.Context, assetSlug string, params map[string]interface{}) (res *AssetMetrics, err error) {
	res = &AssetMetrics{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/assets/%s/metrics", BaseURLV1, assetSlug), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		if fields, ok := params["fields"].([]string); ok {
			q.Add("fields", strings.Join(fields, ","))
		}
	}

	req = req.WithContext(ctx)

	if err := c.fetch(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetAllAssets calls /api/v2/assets
func (c *Client) GetAllAssets(ctx context.Context, params map[string]interface{}) (res []Asset, err error) {
	res = []Asset{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/assets", BaseURLV2), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		if fields, ok := params["fields"].([]string); ok {
			q.Add("fields", strings.Join(fields, ","))
		}
		if limit, ok := params["limit"].(string); ok {
			q.Add("limit", limit)
		}
		if page, ok := params["page"].(string); ok {
			q.Add("page", page)
		}
	}

	req = req.WithContext(ctx)

	if err := c.fetch(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
