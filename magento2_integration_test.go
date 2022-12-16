//go:build integration

package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMagentoBaseURLsFromDatabase(t *testing.T) {
	baseURLs, err := m2store.BaseURLs(fixtureBase + "magento2_integration")
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"https://sansec.io/", "https://api.sansec.io/"}, baseURLs)
}
