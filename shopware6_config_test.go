package gocommerce

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestConfigWithInterpolationToDSN(t *testing.T) {
	sw6 := Shopware6{}
	want := DBConfig{
		Host: "db.internal",
		Port: 3307,
		Name: "caseys_prod",
		User: "caseys",
		Pass: "s3cr3t$$",
	}

	if got := dbConfigFromSource(t, fixtureBase+"/shopware6/configs/.env.interpolated", &sw6); got.DSN() != want.DSN() {
		t.Errorf("ConfigToDSN() = %v, want %v", got.DSN(), want.DSN())
	}
}

func TestFindStoreAtRootShopware6(t *testing.T) {
	root := t.TempDir()
	marker := filepath.Join(root, platformByName(t, "Shopware 6").UniquePath())
	if err := os.MkdirAll(filepath.Dir(marker), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(marker, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	env, err := os.ReadFile(fixtureBase + "/shopware6/configs/.env.interpolated")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, ".env"), env, 0o644); err != nil {
		t.Fatal(err)
	}

	s := FindStoreAtRoot(root, Options{})
	assert.NotNil(t, s)
	assert.Equal(t, "Shopware 6", s.Platform.Name())
	assert.Equal(t, "caseys", s.Config.DB.User)
	assert.Equal(t, 3307, s.Config.DB.Port)
}
