package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var m2store = magento2{
	basePlatform{
		"Magento2",
		"app/etc/env.php",
		"app/etc/env.php",
	},
	"n98-magerun2",
}

func TestM2ConfigSimpleDB(t *testing.T) {
	cfg, err := m2store.ParseConfig(fixtureBase + "magento2/app/etc/env.php")
	assert.NoError(t, err)
	assert.Equal(t, "app:sldfjlskdfklds@tcp(localhost:3306)/magento2?allowOldPasswords=true", cfg.DB.DSN())
	assert.Equal(t, "admin_c6018w", cfg.AdminSlug)
}

func TestVariousM2Configs(t *testing.T) {
	tests := []struct{ path, want string }{
		{"crash1.php", "xx:xx@tcp(localhost:3306)/xx?allowOldPasswords=true"},
		{"crash2.php", "xx:xx@tcp(10.10.20.39:3306)/xx?allowOldPasswords=true"},
		{"crash3.php", "myuser:mypass@tcp(myhost:3306)/mydb?allowOldPasswords=true"},
		{"hostport.php", "gooduser:verylongpassword@tcp(goodhost:3309)/gooddb?allowOldPasswords=true"},
		{"multidb.php", "gooduser:goodpass@tcp(goodhost:3306)/gooddb?allowOldPasswords=true"},
		{"simple.php", "gooduser:verylongpassword@tcp(goodhost:3306)/gooddb?allowOldPasswords=true"},
	}
	for _, test := range tests {
		cfg, err := m2store.ParseConfig(fixtureBase + "magento2/app/etc/" + test.path)
		assert.NoError(t, err)
		assert.Equal(t, test.want, cfg.DB.DSN())
	}
}

func TestEmptyConfig(t *testing.T) {
	cfg, err := m2store.ParseConfig(fixtureBase + "magento2/app/etc/empty.php")
	assert.Error(t, err)
	assert.Nil(t, cfg)
}

func TestGetMagentoVersionFromLockFile(t *testing.T) {
	version, err := m2store.Version(fixtureBase + "magento2")
	assert.Nil(t, err)
	assert.Equal(t, "2.4.2-p2", version)
}

func TestGetMagentoVersionWithoutLockFile(t *testing.T) {
	version, err := m2store.Version(fixtureBase + "magento2_no_lockfile")
	assert.Nil(t, err)
	assert.Equal(t, "2.4.2-p2", version)
}

func TestGetMagentoVersionWithoutSystemPackages(t *testing.T) {
	version, err := m2store.Version(fixtureBase + "magento2_git_clone")
	assert.Nil(t, err)
	assert.Equal(t, "2.4.2-p2", version)
}

func TestGetMagentoBaseURLsFromConfig(t *testing.T) {
	baseURLs, err := m2store.BaseURLs(fixtureBase + "magento2")
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"https://sansec.io/", "https://api.sansec.io/"}, baseURLs)
}
