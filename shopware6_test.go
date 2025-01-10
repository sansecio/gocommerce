package gocommerce

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetShopwareVersionFromLockFile(t *testing.T) {
	ver, err := getVersionFromComposer(filepath.Join(fixtureBase, "shopware6"), sw6ComposerRgx)
	assert.NoError(t, err)
	assert.Equal(t, "6.6.9.0", ver)
}

func TestGetShopwareVersionWithoutLockFile(t *testing.T) {
	ver, err := getVersionFromComposer(filepath.Join(fixtureBase, "shopware6_no_lockfile"), sw6ComposerRgx)
	assert.NoError(t, err)
	assert.Equal(t, "6.6.9.0", ver)
}
