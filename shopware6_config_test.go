package gocommerce

import (
	"testing"
)

var shopware6 = Shopware6{}

func TestConfigToDSN(t *testing.T) {
	want := DBConfig{
		Host: "localhost",
		Port: 3306,
		Name: "db-1",
		User: "db-user-1",
		Pass: "rhPb5xC2242444mFZDB",
	}

	if got := dbConfigFromSource(t, fixtureBase+"/shopware6/configs/.env", &shopware6); got.DSN() != want.DSN() {
		t.Errorf("ConfigToDSN() = %v, want %v", got, want)
	}
}
