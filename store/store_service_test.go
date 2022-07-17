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
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "http://www.baidu.com"
	useruuid := "akdjfa-dfadjl-dat"
	shortURL := "DSKJFuGD"
	SaveUrlMapping(shortURL, initialLink, useruuid)
	retrievedUrl := RetrieveInitialUrl(shortURL)
	assert.Equal(t, initialLink, retrievedUrl)
}
