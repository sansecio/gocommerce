//go:build integration

package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMagento1BaseURLsFromDatabase(t *testing.T) {
	baseURLs, err := m1store.BaseURLs(fixtureBase + "magento1_integration")
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"https://app.magento1.test/", "https://second.magento1.test/"}, baseURLs)
}
