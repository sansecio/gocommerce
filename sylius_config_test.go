package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// .env holds a %kernel.environment% placeholder DSN plus a sibling
// DATABASE_URL_MAGENTO and a commented pgsql line; only the mysql
// DATABASE_URL is picked and %kernel.environment% resolves via APP_ENV=dev.
func TestSyliusConfig(t *testing.T) {
	t.Setenv("DATABASE_URL", "")
	t.Setenv("APP_ENV", "")

	syl := platformByName(t, "Sylius")
	dbc := dbConfigFromSource(t, fixtureBase+"sylius/.env", syl)
	assert.Equal(t,
		"root:@tcp(127.0.0.1:3306)/sylius_dev?allowOldPasswords=true",
		dbc.DSN())
}

// a real DATABASE_URL env var (PaaS injection) overrides the .env file.
func TestSyliusConfigEnvOverride(t *testing.T) {
	t.Setenv("DATABASE_URL", "mysql://u5er:p4ss@db.example.net:10239/randomdb")

	syl := platformByName(t, "Sylius")
	dbc := dbConfigFromSource(t, fixtureBase+"sylius/.env", syl)
	assert.Equal(t,
		"u5er:p4ss@tcp(db.example.net:10239)/randomdb?allowOldPasswords=true",
		dbc.DSN())
}

func TestSyliusVersion(t *testing.T) {
	syl := platformByName(t, "Sylius")
	ver, err := syl.Version(fixtureBase + "sylius")
	assert.NoError(t, err)
	assert.Equal(t, "1.14.4", ver) // anchored regex skips sylius/sylius-rector
}

func TestSyliusDiscovery(t *testing.T) {
	store := FindStoreAtRoot(fixtureBase + "sylius")
	assert.NotNil(t, store)
	assert.Equal(t, "Sylius", store.Platform.Name())
}

func TestSyliusUniquePath(t *testing.T) {
	store := FindStoreAtUniquePath(
		fixtureBase + "sylius/vendor/sylius/sylius/composer.json")
	assert.NotNil(t, store)
	assert.Equal(t, "Sylius", store.Platform.Name())
}
