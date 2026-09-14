package gocommerce

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpandSymfonyVars(t *testing.T) {
	vars := map[string]string{
		"DB_HOST":  "db.internal",
		"DB_PORT":  "3307",
		"PASS":     `pa\$\$word`,
		"INDIRECT": "${DB_HOST}",
		"LOOP":     "${LOOP}",
	}

	for _, tc := range []struct {
		name string
		in   string
		want string
	}{
		{"braces", "mysql://u@${DB_HOST}:${DB_PORT}/db", "mysql://u@db.internal:3307/db"},
		{"bare", "mysql://u@$DB_HOST/db", "mysql://u@db.internal/db"},
		{"nested", "${INDIRECT}", "db.internal"},
		{"default unset", "${DB_SOCKET:-/tmp/mysql.sock}", "/tmp/mysql.sock"},
		{"default set", "${DB_HOST:-localhost}", "db.internal"},
		{"assign default", "${DB_SOCKET:=/run/mysqld.sock}", "/run/mysqld.sock"},
		{"unknown", "mysql://u:${NOPE}@h/db", "mysql://u:@h/db"},
		{"literal from single quotes", "mysql://u:${PASS}@h/db", "mysql://u:pa$$word@h/db"},
		{"escaped dollar", `pa\$sword`, "pa$sword"},
		{"self reference", "${LOOP}", "${LOOP}"},
		{"no reference", "mysql://u:p@h:3306/db", "mysql://u:p@h:3306/db"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, expandSymfonyVars(tc.in, vars, Options{}))
		})
	}
}

func TestExpandSymfonyVarsExplicitEnvironment(t *testing.T) {
	t.Setenv("DB_HOST", "ambient.internal")
	opts := Options{Environment: map[string]string{"DB_HOST": "allowed.internal"}}
	assert.Equal(t, "allowed.internal", expandSymfonyVars("${DB_HOST}", map[string]string{"DB_HOST": "file.internal"}, opts))
}

func TestParseSymfonyDSNEmptyPort(t *testing.T) {
	db := parseSymfonyDSN("DATABASE_URL=mysql://caseys:s3cr3t@db.internal:/caseys_prod", "prod")
	assert.NotNil(t, db)
	assert.Equal(t, 3306, db.Port)
	assert.Equal(t, "caseys_prod", db.Name)
}

func TestSymfonyConfigEnvironmentIsolation(t *testing.T) {
	t.Setenv("ECOMSCAN_KEY", "scanner-license")
	t.Setenv("ZZ_CI_SECRET", "scanner-secret")
	t.Setenv("DATABASE_URL", "mysql://ambient:secret@ambient.internal/ambient")
	t.Setenv("APP_ENV", "ambient")

	for _, name := range []string{"Shopware 6", "Sylius"} {
		t.Run(name, func(t *testing.T) {
			configPath := filepath.Join(t.TempDir(), ".env")
			require.NoError(t, os.WriteFile(configPath, []byte(
				"ECOMSCAN_KEY=file-user\n"+
					"INDIRECT=${ECOMSCAN_KEY}\n"+
					"ECOMSCAN_ALLOW_ENV=ZZ_CI_SECRET\n"+
					"APP_ENV=file\n"+
					`DATABASE_URL="mysql://${INDIRECT}:${ZZ_CI_SECRET:-file-pass}@${ZZ_CI_SECRET:-file-host}/shop_%kernel.environment%"`+"\n"), 0o600))

			cfg, err := platformByName(t, name).ParseConfig(configPath, Options{})
			require.NoError(t, err)
			require.Equal(t, "file-user", cfg.DB.User)
			require.Equal(t, "file-pass", cfg.DB.Pass)
			require.Equal(t, "file-host", cfg.DB.Host)
			require.Equal(t, "shop_file", cfg.DB.Name)
		})
	}
}

func TestSymfonyConfigExplicitEnvironment(t *testing.T) {
	t.Setenv("DB_USER", "ambient-user")
	t.Setenv("ECOMSCAN_KEY", "scanner-license")
	opts := Options{Environment: map[string]string{
		"DB_USER":     "allowed-user",
		"DB_PASSWORD": `pa$ECOMSCAN_KEY\word`,
		"APP_ENV":     "staging",
	}}

	for _, name := range []string{"Shopware 6", "Sylius"} {
		t.Run(name, func(t *testing.T) {
			configPath := filepath.Join(t.TempDir(), ".env")
			require.NoError(t, os.WriteFile(configPath, []byte(
				"DB_USER=file-user\n"+
					"APP_ENV=file\n"+
					`DATABASE_URL="mysql://${DB_USER}:${DB_PASSWORD}@db.internal/shop_%kernel.environment%"`+"\n"), 0o600))

			cfg, err := platformByName(t, name).ParseConfig(configPath, opts)
			require.NoError(t, err)
			require.Equal(t, "allowed-user", cfg.DB.User)
			require.Equal(t, `pa$ECOMSCAN_KEY\word`, cfg.DB.Pass)
			require.Equal(t, "shop_staging", cfg.DB.Name)

			override := Options{Environment: map[string]string{
				"DATABASE_URL": "mysql://override:password@trusted.internal/shop_%kernel.environment%",
				"APP_ENV":      "production",
			}}
			cfg, err = platformByName(t, name).ParseConfig(configPath+".missing", override)
			require.NoError(t, err)
			require.Equal(t, "override", cfg.DB.User)
			require.Equal(t, "trusted.internal", cfg.DB.Host)
			require.Equal(t, "shop_production", cfg.DB.Name)
		})
	}
}
