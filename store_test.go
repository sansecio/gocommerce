package gocommerce

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindStoreAtRoot(t *testing.T) {
	root := fixtureBase + "/magento1"
	s := FindStoreAtRoot(root)
	assert.NotNil(t, s)
	assert.Equal(t, "Magento 1", s.Platform.Name())
}

func TestDiscoverStores(t *testing.T) {
	oldHome := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHome)

	os.Setenv("HOME", fixtureBase+"/discovery")
	stores := DiscoverStores()

	assert.Len(t, stores, 2)

	assert.Equal(t, "Magento 2", stores[0].Platform.Name())
	assert.Contains(t, stores[0].DocRoot, "discovery/public_html/current")

	assert.Equal(t, "WooCommerce", stores[1].Platform.Name())
	assert.Contains(t, stores[1].DocRoot, "discovery/applications/123456789/public_html")
}

func TestDiscoverStoresEmpty(t *testing.T) {
	oldHome := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHome)
	os.Setenv("HOME", "/nonexistant1212123123")
	stores := DiscoverStores()
	assert.Len(t, stores, 0)
	assert.NotNil(t, stores)
}
