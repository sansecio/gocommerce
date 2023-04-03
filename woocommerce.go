package gocommerce

import (
	"os"
	"regexp"
)

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
