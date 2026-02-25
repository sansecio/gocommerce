package gocommerce

import (
	"context"
	"errors"
	"path/filepath"
	"strconv"

	"github.com/sansecio/gocommerce/phpcfg"
)

type Prestashop7 struct {
	basePlatform
}

func (p *Prestashop7) ParseConfig(cfgPath string) (*StoreConfig, error) {
	cm, err := phpcfg.ParsePath(cfgPath)
	if err != nil {
		return nil, err
	}

	port := 3306
	if dbPort := cm["root.parameters.database_port"]; dbPort != "" {
		if pi, err := strconv.Atoi(dbPort); err == nil {
			port = pi
		}
	}

	return &StoreConfig{
		DB: &DBConfig{
			Host:   cm["root.parameters.database_host"],
			User:   cm["root.parameters.database_user"],
			Pass:   cm["root.parameters.database_password"],
			Name:   cm["root.parameters.database_name"],
			Prefix: cm["root.parameters.database_prefix"],
			Port:   port,
		},
	}, nil
}

func (p *Prestashop7) BaseURLs(ctx context.Context, docroot string) ([]string, error) {
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

func (p *Prestashop7) Version(docroot string) (string, error) {
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
