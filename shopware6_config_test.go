package gocommerce

import (
	"testing"
)

func TestConfigToDSN(t *testing.T) {
	sw6 := Shopware6{}
	want := DBConfig{
		Host: "localhost",
		Port: 3306,
		Name: "db-1",
		User: "db-user-1",
		Pass: "rhPb5xC2242444mFZDB",
	}

	if got := dbConfigFromSource(t, fixtureBase+"/shopware6/configs/.env", &sw6); got.DSN() != want.DSN() {
		t.Errorf("ConfigToDSN() = %v, want %v", got, want)
	}
}

func TestConfigWithQuotesToDSN(t *testing.T) {
	sw6 := Shopware6{}
	want := DBConfig{
		Host: "localhost",
		Port: 3306,
		Name: "DB",
		User: "USER",
		Pass: "PASS",
	}

	if got := dbConfigFromSource(t, fixtureBase+"/shopware6/configs/.env.quotes", &sw6); got.DSN() != want.DSN() {
		t.Errorf("ConfigToDSN() = %v, want %v", got.DSN(), want.DSN())
	}
}
