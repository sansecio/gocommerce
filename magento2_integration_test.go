//go:build integration

package gocommerce

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMagento2BaseURLsFromDatabase(t *testing.T) {
	baseURLs, err := m2store.BaseURLs(context.TODO(), fixtureBase+"magento2_integration", Options{})
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"https://sansec.io/", "https://api.sansec.io/"}, baseURLs)
}
