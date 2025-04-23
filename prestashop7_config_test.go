package gocommerce

import (
	"testing"

	"gotest.tools/assert"
)

func TestPrestashop7Config(t *testing.T) {
	dbc := dbConfigFromSource(t, fixtureBase+"/prestashop7/configs/normal.php", &Prestashop7{})
	assert.Equal(t, "presta_user:qwerty123@tcp(127.0.0.1:3306)/presta_db?allowOldPasswords=true", dbc.DSN())
}

func TestPrestashop7ConfigWithCustomPort(t *testing.T) {
	dbc := dbConfigFromSource(t, fixtureBase+"/prestashop7/configs/custom-port.php", &Prestashop7{})
	assert.Equal(t, "presta_user:qwerty123@tcp(127.0.0.1:1234)/presta_db?allowOldPasswords=true", dbc.DSN())
}
