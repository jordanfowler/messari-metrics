package main

import (
	"context"
	"fmt"
	"os"

	messari "github.com/jordanfowler/messari-metrics/messari"
)

const (
	assetsLimit = 500
)

// AssetCaches keeps fast look up of a cache entry by key/value pair
var AssetCaches = map[string]AssetCache{}

// AssetCache groups assets by common key/value
type AssetCache map[string]bool

// PushAssetCache upserts asset slug into cache by key/value pair
func PushAssetCache(key, value, assetSlug string) bool {
	cacheKey := fmt.Sprintf("%s-%s", key, value)
	if _, ok := AssetCaches[cacheKey]; !ok {
		AssetCaches[cacheKey] = AssetCache{}
	}
	AssetCaches[cacheKey][assetSlug] = true
	return AssetCaches[cacheKey][assetSlug]
}

// GetAssetCacheSlugs returns the asset slugs for a given key/value pair
func GetAssetCacheSlugs(key, value string) (assetSlugs []string) {
	cacheKey := fmt.Sprintf("%s-%s", key, value)
	if _, ok := AssetCaches[cacheKey]; !ok {
		return
	}
	for assetSlug := range AssetCaches[cacheKey] {
		assetSlugs = append(assetSlugs, assetSlug)
	}
	return
}

// buildAssetCaches will load all assets and cache their tags and sector data for faster aggregate responses (reduce number of parallel asset requests)
func buildAssetCaches() {
	ctx := context.Background()
	client := messari.NewClient(os.Getenv("MESSARI_API_KEY"))

	var assetsPage = 1
	for {
		fmt.Println(fmt.Sprintf("buildAssetCaches loading page=%v limit=%v", assetsPage, assetsLimit))

		assets, err := client.GetAllAssets(ctx, map[string]interface{}{
			"page":  fmt.Sprint(assetsPage),
			"limit": fmt.Sprint(assetsLimit),
			"fields": []string{
				"market_data",
			},
		})

		if err != nil {
			fmt.Println(fmt.Sprintf("Error [GetAllAssets]: %s", err))
			break
		}

		for _, asset := range assets {
			for _, tag := range asset.Metrics.MiscData.Tags {
				PushAssetCache("tags", tag, asset.Slug)
			}
			for _, sector := range asset.Metrics.MiscData.Sector {
				PushAssetCache("sectors", sector, asset.Slug)
			}
		}

		if len(assets) == 0 {
			break
		} else {
			assetsPage++
		}
		fmt.Println(AssetCaches)
		break
	}

}
