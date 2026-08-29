package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestExpandSymfonyVarsPrefersRealEnv(t *testing.T) {
	t.Setenv("DB_HOST", "10.0.0.9")
	assert.Equal(t, "10.0.0.9", expandSymfonyVars("${DB_HOST}", map[string]string{"DB_HOST": "db.internal"}))
}

func TestParseSymfonyDSNEmptyPort(t *testing.T) {
	db := parseSymfonyDSN("DATABASE_URL=mysql://caseys:s3cr3t@db.internal:/caseys_prod", "prod")
	assert.NotNil(t, db)
	assert.Equal(t, 3306, db.Port)
	assert.Equal(t, "caseys_prod", db.Name)
}
