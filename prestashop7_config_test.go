package gocommerce

import (
	"testing"

	"gotest.tools/assert"
)

func TestPrestashopConfig(t *testing.T) {
	p7 := Prestashop7{}
	dbc := dbConfigFromSource(t, fixtureBase+"/prestashop7/configs/normal.php", &p7)
	assert.Equal(t, "presta_user:qwerty123@tcp(127.0.0.1:3306)/presta_db?allowOldPasswords=true", dbc.DSN())
}
