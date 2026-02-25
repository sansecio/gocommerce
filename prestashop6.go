package gocommerce

import (
	"context"
	"errors"
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

	db, err := ConnectDB(ctx, *cfg.DB)
	if err != nil {
		return nil, err
	}

	prefix, err := cfg.DB.SafePrefix()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT DISTINCT domain_ssl FROM ` + prefix + `shop_url`)
	if err != nil {
		return nil, err
	}

	urls := []string{}
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err == nil {
			urls = append(urls, "https://"+url+"/")
		}
	}

	if len(urls) > 0 {
		return urls, nil
	}

	return nil, errors.New("base url(s) not found in database")
}

func (p *Prestashop6) Version(docroot string) (string, error) {
	cfg, err := p.ParseConfig(filepath.Join(docroot, p.ConfigPath()))
	if err != nil {
		return "", err
	}

	db, err := ConnectDB(context.Background(), *cfg.DB)
	if err != nil {
		return "", err
	}

	prefix, err := cfg.DB.SafePrefix()
	if err != nil {
		return "", err
	}

	var version string
	err = db.QueryRow(`SELECT value FROM ` + prefix + `configuration WHERE name = 'PS_VERSION_DB'`).Scan(&version)
	if err != nil {
		return "", err
	}

	return version, nil
}
