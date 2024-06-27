package gocommerce

import (
	"testing"

	"gotest.tools/assert"
)

func TestPrestashop6Config(t *testing.T) {
	dbc := dbConfigFromSource(t, fixtureBase+"/prestashop6/config/settings.inc.php", &Prestashop6{})
	assert.Equal(t, "sansec_user:sansec_password@tcp(localhost:1234)/sansec_db?allowOldPasswords=true", dbc.DSN())
}
