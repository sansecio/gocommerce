package gocommerce

import (
	"testing"

	"gotest.tools/assert"
)

var prestashop7 = Prestashop7{}

func TestNormal(t *testing.T) {
	dbc := dbConfigFromSource(t, fixtureBase+"/prestashop7/configs/normal.php", &prestashop7)
	assert.Equal(t, "presta_user:qwerty123@tcp(127.0.0.1:3306)/presta_db?allowOldPasswords=true", dbc.DSN())
}
