package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.NotNil(t, testStoreService.redisClient, "Redis client should not be nil after initialization")
}

func testInsertAndRetrieve(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	SaveUrlMapping(shortURL, initialLink, userUUId)

	retrievedUrl := RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrievedUrl, "The retrieved URL should match the original URL")
	
}