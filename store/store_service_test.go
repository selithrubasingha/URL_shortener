package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


// IMPORTANT : A test is just a script that feeds fake inputs into your
//  functions and checks if the output matches what it expects.

// go test functions in a file run top to bottom by default . but you can change it if you want . 

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {

	// assert == CHECK IF ... so assert.NotNil == CHECK IF NOT NIL
	assert.NotNil(t, testStoreService.redisClient, "Redis client should not be nil after initialization")
}

//testing methods must ALWAYS start with "Test".
// also the file name should end with "_test.go" so that the go test command can find it and run it.
func TestInsertAndRetrieve(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	SaveUrlMapping(shortURL, initialLink, userUUId)

	retrievedUrl := RetrieveInitialUrl(shortURL)

	// assert.Equal checks if the two values are equal and if not, it will fail the test and print the message provided
	assert.Equal(t, initialLink, retrievedUrl, "The retrieved URL should match the original URL")
	
}