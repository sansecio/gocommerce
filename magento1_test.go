package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestM1ConfigSimpleDB(t *testing.T) {
	src := fixtureBase + "magento1/local.xml"
	cfg, err := ParseMagento1Config(src)
	assert.NoError(t, err)
	assert.Equal(t, "app:sdkjfhjksdfjk@tcp(localhost:3306)/mag1?allowOldPasswords=true", cfg.DSN())
	assert.Equal(t, "s3cr3tfrontname", cfg.AdminSlug)
}

func TestM1ConfigHostWithPort(t *testing.T) {
	src := fixtureBase + "magento1/local.xml.hostport"
	cfg, err := ParseMagento1Config(src)
	assert.NoError(t, err)
	assert.Equal(t, "app:werrrwww@tcp(localhost:3307)/mag1?allowOldPasswords=true", cfg.DSN())
}

func TestM1EmptyPassword(t *testing.T) {
	src := fixtureBase + "magento1/local.xml.nopass"
	cfg, err := ParseMagento1Config(src)
	assert.NoError(t, err)
	assert.Equal(t, "userhere:@tcp(db:3306)/dbnamehere?allowOldPasswords=true", cfg.DSN())
}

func TestM1BogusConfig(t *testing.T) {
	src := fixtureBase + "magento1/local.xml.bogus"
	cfg, err := ParseMagento1Config(src)
	assert.Nil(t, cfg)
	assert.Error(t, err)
}
