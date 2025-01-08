package gocommerce

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"regexp"
)

type WooCommerce struct {
	basePlatform
}

var wpLookupRgx = map[string]string{
	"user":   `define\(\s*['"]DB_USER['"]\s*,\s*['"](\S+?)['"]\s*\);`,
	"pass":   `define\(\s*['"]DB_PASSWORD['"]\s*,\s*['"]([^']{0,64})['"]\s*\);`,
	"host":   `define\(\s*['"]DB_HOST['"]\s*,\s*['"](\S+?)['"]\s*\);`,
	"db":     `define\(\s*['"]DB_NAME['"]\s*,\s*['"](\S+?)['"]\s*\);`,
	"prefix": `$?table_prefix\s*=\s*['"]([^']*?)['"]\s*;`,
}

func (w *WooCommerce) ParseConfig(cfgPath string) (*StoreConfig, error) {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	matches := map[string]string{}
	port := 3306

	for k, v := range wpLookupRgx {
		m := regexp.MustCompile(v).FindStringSubmatch(string(data))
		if len(m) != 2 {
			continue
		}
		matches[k] = m[1]
	}

	if h, p, e := parseHostPort(matches["host"]); p > 0 && e == nil {
		matches["host"] = h
		port = p
	}

	return &StoreConfig{
		DB: &DBConfig{
			Host:   matches["host"],
			User:   matches["user"],
			Pass:   matches["pass"],
			Name:   matches["db"],
			Prefix: matches["prefix"],
			Port:   port,
		},
	}, nil
}

func (w *WooCommerce) BaseURLs(ctx context.Context, docroot string) ([]string, error) {
	cfg, err := w.ParseConfig(filepath.Join(docroot, w.ConfigPath()))
	if err != nil {
		return nil, err
	}

	db, err := ConnectDB(ctx, *cfg.DB)
	if err != nil {
		return nil, err
	}

	prefix, err := cfg.DB.SafePrefix()
	if err != nil {
		return nil, err
	}

	var url string
	if err = db.QueryRow(`select option_value from ` + prefix + `options where option_name = 'home'`).Scan(&url); err != nil {
		return nil, err
	}
	return []string{url}, nil
}

func (w *WooCommerce) Version(docroot string) (string, error) {
	re := regexp.MustCompile(`\$wp_version\s*=\s*'([^']+)';`)
	data, err := os.ReadFile(filepath.Join(docroot, "wp-includes", "version.php"))
	if err != nil {
		return "", err
	}
	match := re.FindStringSubmatch(string(data))
	if len(match) < 2 {
		return "", errors.New("no version found")
	}
	return match[1], nil
}
