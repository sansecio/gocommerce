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

// Shared helpers for Symfony-based platforms (Shopware 6, Sylius, ...).
// They read the database connection from a Symfony DATABASE_URL, either from a
// real environment variable or from .env file.

var (
	// matches a mysql DATABASE_URL, ignoring commented lines and the
	// DATABASE_URL_<SUFFIX> variants, with optional password, optional port
	// and optional ?query parameters. The DATABASE_URL= prefix is optional so
	// a bare DSN (from the environment) parses with the same expression.
	symfonyDBURLRgx = regexp.MustCompile(
		`(?m)^\s*(?:DATABASE_URL\s*=\s*)?"?(?:pdo-)?mysql://([^:@/]+)(?::([^@]*))?@([^:/?]+)(?::(\d+))?/([^?"\s]+)`)

	symfonyAppEnvRgx = regexp.MustCompile(`(?m)^\s*APP_ENV\s*=\s*"?([a-zA-Z0-9_-]+)`)
)

// symfonyParseConfig parses a Symfony DATABASE_URL into a StoreConfig. A real
// DATABASE_URL environment variable takes precedence over the .env file,
// mirroring Symfony's own dotenv override behaviour.
func symfonyParseConfig(cfgPath string) (*StoreConfig, error) {
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		if db := parseSymfonyDSN(dsn, kernelEnv("")); db != nil {
			return &StoreConfig{DB: db}, nil
		}
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	db := parseSymfonyDSN(string(data), kernelEnv(string(data)))
	if db == nil {
		return nil, fmt.Errorf("could not parse mysql DATABASE_URL in %s", cfgPath)
	}
	return &StoreConfig{DB: db}, nil
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

// kernelEnv resolves the value Symfony substitutes for %kernel.environment%:
// a real APP_ENV environment variable wins, else the APP_ENV from the .env
// contents, else "prod" as a production-scan default.
func kernelEnv(fileData string) string {
	if v := os.Getenv("APP_ENV"); v != "" {
		return v
	}
	if m := symfonyAppEnvRgx.FindStringSubmatch(fileData); m != nil {
		return m[1]
	}
	return "prod"
}

// symfonyColumnURLs runs a single-column query and collects the non-empty
// results, the shared mechanics behind every Symfony platform's BaseURLs.
func symfonyColumnURLs(ctx context.Context, cfg *StoreConfig, query string) ([]string, error) {
	db, err := ConnectDB(ctx, *cfg.DB)
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
