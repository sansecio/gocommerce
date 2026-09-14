package gocommerce

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindStoreAtRoot(t *testing.T) {
	root := fixtureBase + "/magento1"
	s := FindStoreAtRoot(root, Options{})
	assert.NotNil(t, s)
	assert.Equal(t, "Magento 1", s.Platform.Name())
}

func TestDiscoverStores(t *testing.T) {
	oldHome := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHome)

	os.Setenv("HOME", fixtureBase+"/discovery")
	stores := DiscoverStores(Options{})

	assert.Len(t, stores, 2)

	assert.Equal(t, "WP/WooCommerce", stores[0].Platform.Name())
	assert.Contains(t, stores[0].DocRoot, "discovery/applications/123456789/public_html")

	assert.Equal(t, "Magento 2", stores[1].Platform.Name())
	assert.Contains(t, stores[1].DocRoot, "discovery/sansec.store")
}

func TestDiscoverSymlinkedStore(t *testing.T) {
	os.Setenv("HOME", fixtureBase+"/discovery/symlink")
	stores := DiscoverStores(Options{})

	assert.Len(t, stores, 1)
	assert.Equal(t, "Magento 2", stores[0].Platform.Name())
	assert.Contains(t, stores[0].DocRoot, "discovery/sansec.store")
}

// an unreadable config must not hide the platform, callers keep every check
// that does not need the database, and reading the empty config is safe.
func TestFindStoreAtRootUnreadableConfig(t *testing.T) {
	root := t.TempDir()
	// a directory named wp-config.php makes os.ReadFile fail while os.Stat succeeds
	if err := os.Mkdir(root+"/wp-config.php", 0o755); err != nil {
		t.Fatal(err)
	}

	s := FindStoreAtRoot(root, Options{})
	assert.NotNil(t, s)
	assert.Equal(t, "WP/WooCommerce", s.Platform.Name())
	assert.Empty(t, s.Config.DB.Prefix)
	assert.Empty(t, s.Config.DB.DSN())
}

func TestDiscoverStoresEmpty(t *testing.T) {
	oldHome := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHome)
	os.Setenv("HOME", "/nonexistant1212123123")
	stores := DiscoverStores(Options{})
	assert.Len(t, stores, 0)
	assert.NotNil(t, stores)
}
