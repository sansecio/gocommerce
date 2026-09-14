package gocommerce

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	// matches a mysql DATABASE_URL, ignoring commented lines and the
	// DATABASE_URL_<SUFFIX> variants, with optional password, optional port
	// and optional ?query parameters. The DATABASE_URL= prefix is optional so
	// a bare DSN (from the environment) parses with the same expression.
	symfonyDBURLRgx = regexp.MustCompile(
		`(?m)^\s*(?:DATABASE_URL\s*=\s*)?"?(?:pdo-)?mysql://([^:@/]+)(?::([^@]*))?@([^:/?]+)(?::(\d*))?/([^?"\s]+)`)

	symfonyAppEnvRgx = regexp.MustCompile(`(?m)^\s*APP_ENV\s*=\s*"?([a-zA-Z0-9_-]+)`)

	// a KEY=VALUE assignment in a .env file, with optional `export ` prefix
	symfonyEnvLineRgx = regexp.MustCompile(`(?m)^[ \t]*(?:export[ \t]+)?([A-Za-z_][A-Za-z0-9_]*)[ \t]*=[ \t]*(.*)$`)

	// a variable reference inside a .env value: $VAR, ${VAR} or ${VAR:-default}
	symfonyVarRefRgx = regexp.MustCompile(`\$(?:\{([A-Za-z_][A-Za-z0-9_]*)(?::[-=]([^}]*))?\}|([A-Za-z_][A-Za-z0-9_]*))`)
)

func symfonyParseConfig(cfgPath string, opts Options) (*StoreConfig, error) {
	if dsn := opts.Environment["DATABASE_URL"]; dsn != "" {
		if db := parseSymfonyDSN(dsn, kernelEnv("", opts)); db != nil {
			return &StoreConfig{DB: db}, nil
		}
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	vars := symfonyEnvVars(string(data))
	dsn, ok := vars["DATABASE_URL"]
	if !ok {
		dsn = string(data)
	}

	db := parseSymfonyDSN(expandSymfonyVars(dsn, vars, opts), kernelEnv(string(data), opts))
	if db == nil {
		return nil, fmt.Errorf("could not parse mysql DATABASE_URL in %s", cfgPath)
	}
	return &StoreConfig{DB: db}, nil
}

// symfonyEnvVars collects the KEY=VALUE assignments of a .env file, unquoted,
// with the last assignment of a key winning as in Symfony's own dotenv
func symfonyEnvVars(data string) map[string]string {
	m := symfonyEnvLineRgx.FindAllStringSubmatch(data, -1)
	vars := make(map[string]string, len(m))
	for _, v := range m {
		val := strings.TrimRight(v[2], " \t\r")
		if lit, ok := unquoteSymfonyValue(val, '\''); ok {
			vars[v[1]] = strings.ReplaceAll(lit, "$", `\$`)
			continue
		}
		if val2, ok := unquoteSymfonyValue(val, '"'); ok {
			vars[v[1]] = val2
			continue
		}

		if i := strings.Index(val, " #"); i >= 0 {
			val = strings.TrimRight(val[:i], " \t")
		}
		vars[v[1]] = val
	}
	return vars
}

func unquoteSymfonyValue(v string, q byte) (string, bool) {
	if len(v) < 2 || v[0] != q {
		return v, false
	}
	end := strings.LastIndexByte(v, q)
	if end == 0 {
		return v, false
	}
	return v[1:end], true
}

func expandSymfonyVars(s string, vars map[string]string, opts Options) string {
	const escaped = "\x00" // shields a literal dollar from every expansion round
	for range 8 {
		// also shields the literals that this round's substitutions brought in
		s = strings.ReplaceAll(s, `\$`, escaped)
		if !strings.ContainsRune(s, '$') {
			break
		}
		expanded := symfonyVarRefRgx.ReplaceAllStringFunc(s, func(ref string) string {
			m := symfonyVarRefRgx.FindStringSubmatch(ref)
			name, fallback := m[1], m[2]
			if name == "" {
				name = m[3]
			}
			if v := opts.Environment[name]; v != "" {
				return strings.ReplaceAll(v, "$", escaped)
			}
			if v := vars[name]; v != "" {
				return v
			}
			return fallback
		})
		if expanded == s {
			break
		}
		s = expanded
	}
	return strings.ReplaceAll(s, escaped, "$")
}

// parseSymfonyDSN extracts mysql connection details from a string holding a
// DATABASE_URL assignment or a bare DSN. env, when set, replaces Symfony's
// %kernel.environment% placeholder in the database name.
func parseSymfonyDSN(s, env string) *DBConfig {
	m := symfonyDBURLRgx.FindStringSubmatch(s)
	if m == nil {
		return nil
	}

	port := 3306
	if m[4] != "" {
		if p, err := strconv.Atoi(m[4]); err == nil {
			port = p
		}
	}

	name := strings.ReplaceAll(m[5], "%kernel.environment%", env)

	return &DBConfig{
		Host: m[3],
		User: m[1],
		Pass: m[2],
		Name: name,
		Port: port,
	}
}

func kernelEnv(fileData string, opts Options) string {
	if v := opts.Environment["APP_ENV"]; v != "" {
		return v
	}
	if m := symfonyAppEnvRgx.FindStringSubmatch(fileData); m != nil {
		return m[1]
	}
	return "prod"
}

// symfonyColumnURLs runs a single-column query and collects the non-empty
// results, the shared mechanics behind every Symfony platform's BaseURLs.
func symfonyColumnURLs(ctx context.Context, cfg *StoreConfig, query string, opts Options) ([]string, error) {
	db, err := ConnectDB(ctx, *cfg.DB, opts)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	urls := []string{}
	for rows.Next() {
		var u string
		if err := rows.Scan(&u); err == nil && u != "" {
			urls = append(urls, u)
		}
	}

	if len(urls) == 0 {
		return nil, errors.New("base url(s) not found in database")
	}
	return urls, nil
}
