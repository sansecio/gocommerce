package gocommerce

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Prestashop6 struct {
	basePlatform
}

var p6LookupRgx = map[string]string{
	"host":   `define\('_DB_SERVER_',\s?'([^']+)'\);`,
	"user":   `define\('_DB_USER_',\s?'([^']+)'\);`,
	"pass":   `define\('_DB_PASSWD_',\s?'([^']+)'\);`,
	"db":     `define\('_DB_NAME_\',\s?'([^']+)'\);`,
	"prefix": `define\('_DB_PREFIX_\',\s?'([^']+)'\);`,
}

func (p *Prestashop6) ParseConfig(cfgPath string) (*StoreConfig, error) {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	if strings.Contains(string(data), "@deprecated") {
		return nil, fmt.Errorf("deprecated config, skipping")
	}

	matches := map[string]string{}
	for k, v := range p6LookupRgx {
		m := regexp.MustCompile(v).FindStringSubmatch(string(data))
		if len(m) != 2 {
			continue
		}
		matches[k] = m[1]
	}

	port := 3306
	if strings.Contains(matches["host"], ":") {
		parts := strings.Split(matches["host"], ":")
		matches["host"] = parts[0]

		newPort, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		port = newPort
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

func (p *Prestashop6) BaseURLs(ctx context.Context, docroot string) ([]string, error) {
	cfg, err := p.ParseConfig(filepath.Join(docroot, p.ConfigPath()))
	if err != nil {
		return nil, err
	}
	return prestashopBaseURLs(ctx, cfg)
}

func (p *Prestashop6) Version(docroot string) (string, error) {
	cfg, err := p.ParseConfig(filepath.Join(docroot, p.ConfigPath()))
	if err != nil {
		return "", err
	}
	return prestashopVersion(cfg)
}
