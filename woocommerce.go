package gocommerce

import (
	"os"
	"regexp"
)

var lookupRgx = map[string]string{
	"user":   `define\(\s*['"]DB_USER['"]\s*,\s*['"](\S+?)['"]\s*\);`,
	"pass":   `define\(\s*['"]DB_PASSWORD['"]\s*,\s*['"]([^']{0,64})['"]\s*\);`,
	"host":   `define\(\s*['"]DB_HOST['"]\s*,\s*['"](\S+?)['"]\s*\);`,
	"db":     `define\(\s*['"]DB_NAME['"]\s*,\s*['"](\S+?)['"]\s*\);`,
	"prefix": `$?table_prefix\s*=\s*['"]([^']*?)['"]\s*;`,
}

func (w *woocommerce) TableChecks() []TableCheck {
	return []TableCheck{
		{"posts", "ID", "post_content", "post_type != 'shop_order' AND length(post_content) > 64"}, // somewhat arbitrary number to kill small files
		{"options", "option_id", "option_value", "length(option_value) > 64"},
		// Temp disabled, silver has 280K users which would take multiple hours to scan
		// There seems no simple way to filter for admin accounts
		// {"users", "user_login", "user_email", ""},
	}
}

func (w *woocommerce) ParseConfig(cfgPath string) (*StoreConfig, error) {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	matches := map[string]string{}

	for k, v := range lookupRgx {
		m := regexp.MustCompile(v).FindStringSubmatch(string(data))
		if len(m) != 2 {
			continue
		}
		matches[k] = m[1]
	}

	return &StoreConfig{
		DB: &DBConfig{
			Host:   matches["host"],
			User:   matches["user"],
			Pass:   matches["pass"],
			Name:   matches["db"],
			Prefix: matches["prefix"],
		},
	}, nil
}
