package gocommerce

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJTLShopConfig(t *testing.T) {
	jtl := platformByName(t, "JTL-Shop")
	dbc := dbConfigFromSource(t, fixtureBase+"jtlshop/includes/config.JTL-Shop.ini.php", jtl)
	assert.Equal(t, "jtl:jtlpass@tcp(localhost:3306)/jtl?allowOldPasswords=true", dbc.DSN())
}

func TestJTLShopURL(t *testing.T) {
	jtl := platformByName(t, "JTL-Shop")
	urls, err := jtl.BaseURLs(context.TODO(), fixtureBase+"jtlshop")
	assert.NoError(t, err)
	assert.Equal(t, []string{"http://sansec.io"}, urls)
}

func TestJTLShopVersion(t *testing.T) {
	jtl := platformByName(t, "JTL-Shop")
	ver, err := jtl.Version(fixtureBase + "jtlshop")
	assert.NoError(t, err)
	assert.Equal(t, "5.5.2", ver)
}
