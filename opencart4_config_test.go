package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenCartConfig(t *testing.T) {
	oc4 := OpenCart4{}
	dbc := dbConfigFromSource(t, fixtureBase+"opencart4/config.php", &oc4)
	assert.Equal(t, dbc.DSN(), "root:root@tcp(localhost:3306)/opencart?allowOldPasswords=true")
}

func TestOpenCartURL(t *testing.T) {
	oc4 := OpenCart4{}
	urls, err := oc4.BaseURLs(fixtureBase + "opencart4", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, urls)
	assert.Equal(t, "http://sansec.io/", urls[0])
}

func TestOpenCartVersion(t *testing.T) {
	oc4 := OpenCart4{}
	ver, err := oc4.Version(fixtureBase + "opencart4")
	assert.NoError(t, err)
	assert.Equal(t, "4.0.2.3", ver)
}
