package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var m1store = Magento1{
	basePlatform{
		"Magento 1",
		"app/etc/local.xml",
		"app/etc/local.xml",
	},
	"n98-magerun2",
}

func TestM1Configs(t *testing.T) {
	tests := []struct{ path, want, slug string }{
		{"local.xml", "app:sdkjfhjksdfjk@tcp(localhost:3306)/mag1?allowOldPasswords=true", "s3cr3tfrontname"},
		{"local.xml.hostport", "app:werrrwww@tcp(localhost:3307)/mag1?allowOldPasswords=true", "willem"},
		{"local.xml.nopass", "userhere:@tcp(db:3306)/dbnamehere?allowOldPasswords=true", "admin"},
	}

	for _, test := range tests {
		src := fixtureBase + "magento1/app/etc/" + test.path
		cfg, err := m1store.ParseConfig(src)
		assert.NoError(t, err)
		assert.Equal(t, test.want, cfg.DB.DSN())
		assert.Equal(t, test.slug, cfg.AdminSlug)
	}
}

func TestM1BogusConfig(t *testing.T) {
	cfg, err := m1store.ParseConfig(fixtureBase + "magento1/app/etc/local.xml.bogus")
	assert.Error(t, err)
	assert.Nil(t, cfg)
}
