package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	assert.Equal(t, shortLink_1, "d66yfx7N")
}
