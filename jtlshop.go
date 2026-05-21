package gocommerce

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type JTLShop struct {
	basePlatform
}

var (
	jtlURLRgx = regexp.MustCompile(`define\(\s*['"]URL_SHOP['"]\s*,\s*['"]([^'"]+)['"]\s*\)`)
	jtlVerRgx = regexp.MustCompile(`APPLICATION_VERSION\s*=\s*['"]([^'"]+)['"]`)

	jtlLookupRgx = map[string]string{
		"host": `define\(\s*['"]DB_HOST['"]\s*,\s*['"]([^'"]+)['"]\s*\)`,
		"user": `define\(\s*['"]DB_USER['"]\s*,\s*['"]([^'"]+)['"]\s*\)`,
		"pass": `define\(\s*['"]DB_PASS['"]\s*,\s*['"]([^'"]{0,128})['"]\s*\)`,
		"db":   `define\(\s*['"]DB_NAME['"]\s*,\s*['"]([^'"]+)['"]\s*\)`,
		"port": `define\(\s*['"]DB_PORT['"]\s*,\s*['"]?(\d+)['"]?\s*\)`,
	}
)

func (j *JTLShop) ParseConfig(cfgPath string) (*StoreConfig, error) {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	matches := map[string]string{}
	for k, v := range jtlLookupRgx {
		m := regexp.MustCompile(v).FindStringSubmatch(string(data))
		if len(m) != 2 {
			continue
		}
		matches[k] = m[1]
	}

	// JTL-Shop only defines DB_PORT when it deviates from the default.
	port := 3306
	if p, err := strconv.Atoi(matches["port"]); err == nil && p > 0 {
		port = p
	}

	return &StoreConfig{
		DB: &DBConfig{
			Host: matches["host"],
			User: matches["user"],
			Pass: matches["pass"],
			Name: matches["db"],
			Port: port,
		},
	}, nil
}

func (j *JTLShop) BaseURLs(_ context.Context, docroot string) ([]string, error) {
	data, err := os.ReadFile(filepath.Join(docroot, j.ConfigPath()))
	if err != nil {
		return nil, err
	}

	m := jtlURLRgx.FindStringSubmatch(string(data))
	if len(m) < 2 {
		return nil, errors.New("base url not found in config")
	}

	return []string{m[1]}, nil
}

func (j *JTLShop) Version(docroot string) (string, error) {
	data, err := os.ReadFile(filepath.Join(docroot, "includes", "defines_inc.php"))
	if err != nil {
		return "", err
	}

	m := jtlVerRgx.FindStringSubmatch(string(data))
	if len(m) < 2 {
		return "", errors.New("version not found in config")
	}

	return m[1], nil
}
