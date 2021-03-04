package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushAssetCache(t *testing.T) {
	assert.True(t, PushAssetCache("tags", "Stablecoins", "btc"))
}

func TestGetAssetCacheSlugs(t *testing.T) {
	assert.Contains(t, GetAssetCacheSlugs("tags", "Stablecoins"), "btc")
}
