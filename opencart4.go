package gocommerce

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type OpenCart4 struct {
	basePlatform
}

var (
	ocURLRgx    = regexp.MustCompile(`define\('HTTP_SERVER',\s?'([^']+)'\);`)
	ocVerRgx    = regexp.MustCompile(`define\('VERSION',\s?'([^']+)'\);`)
	ocLookupRgx = map[string]string{
		"host":   `define\('DB_HOSTNAME\',\s?'([^']+)'\);`,
		"user":   `define\('DB_USERNAME\',\s?'([^']+)'\);`,
		"pass":   `define\('DB_PASSWORD\',\s?'([^']+)'\);`,
		"db":     `define\('DB_DATABASE\',\s?'([^']+)'\);`,
		"port":   `define\('DB_PORT\',\s?'([^']+)'\);`,
		"prefix": `define\('DB_PREFIX\',\s?'([^']+)'\);`,
	}
)

func (oc4 *OpenCart4) ParseConfig(cfgPath string) (*StoreConfig, error) {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	matches := map[string]string{}
	for k, v := range ocLookupRgx {
		m := regexp.MustCompile(v).FindStringSubmatch(string(data))
		if len(m) != 2 {
			continue
		}
		matches[k] = m[1]
	}

	port, err := strconv.Atoi(matches["port"])
	if err != nil {
		return nil, err
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

func (oc4 *OpenCart4) BaseURLs(ctx context.Context, docroot string) ([]string, error) {
	cfgPath := filepath.Join(docroot, "config.php")
	cfg, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	match := ocURLRgx.FindSubmatch(cfg)
	if len(match) < 2 {
		return nil, errors.New("base url not found in config")
	}

	return []string{string(match[1])}, nil
}

func (oc4 *OpenCart4) Version(docroot string) (string, error) {
	cfgPath := filepath.Join(docroot, "admin", "index.php")
	cfg, err := os.ReadFile(cfgPath)
	if err != nil {
		return "", err
	}

	match := ocVerRgx.FindSubmatch(cfg)
	if len(match) < 2 {
		return "", errors.New("version not found in config")
	}

	return string(match[1]), nil
}
