//go:build integration

package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMagento2BaseURLsFromDatabase(t *testing.T) {
	baseURLs, err := m2store.BaseURLs(nil, fixtureBase + "magento2_integration")
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"https://sansec.io/", "https://api.sansec.io/"}, baseURLs)
}
