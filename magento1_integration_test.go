//go:build integration

package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMagentoBaseURLsFromDatabase(t *testing.T) {
	baseURLs, err := m2store.BaseURLs(fixtureBase + "magento1_integration")
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"https://app.magento1.test/", "https://second.magento1.test/"}, baseURLs)
}
