package messari

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAssetMetrics(t *testing.T) {
	c := NewClient(os.Getenv("MESSARI_API_KEY"))
	ctx := context.Background()

	params := map[string]interface{}{
		"fields": []string{"market_data", "marketcap"},
	}
	res, err := c.GetAssetMetrics(ctx, "btc", params)
	t.Logf("res: %v", res)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetAllAssets(t *testing.T) {
	c := NewClient(os.Getenv("MESSARI_API_KEY"))
	ctx := context.Background()

	params := map[string]interface{}{
		"limit": 30,
		"page":  1,
	}
	res, err := c.GetAllAssets(ctx, params)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Len(t, res, 30, "expecting there to be 30 assets")
}
