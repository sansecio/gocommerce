//go:build integration

package gocommerce

import (
	"testing"
	"context"

	"github.com/stretchr/testify/assert"
)

func TestGetMagento2BaseURLsFromDatabase(t *testing.T) {
	baseURLs, err := m2store.BaseURLs(context.TODO(), fixtureBase+"magento2_integration")
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"https://sansec.io/", "https://api.sansec.io/"}, baseURLs)
}
