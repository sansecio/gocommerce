package gocommerce

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type Shopware6 struct {
	basePlatform
}

const (
	// DATABASE_URL=mysql://db-user-1:rhPb5xC2242444mFZDB@localhost:3306/db-1
	DBURL = `(?m)^\s*DATABASE_URL\s*="?\s*mysql://(.+?):(.+?)@(.+?):(\d+)/(.+?)"?$`
)

var sw6ComposerRgx = regexp.MustCompile(`shopware\/core`)

func (s *Shopware6) ParseConfig(cfgPath string) (*StoreConfig, error) {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	m := regexp.MustCompile(DBURL).FindStringSubmatch(string(data))
	if len(m) != 6 {
		return nil, fmt.Errorf("could not parse DSN in Shopware config path %s", cfgPath)
	}

	port, err := strconv.Atoi(m[4])
	if err != nil {
		return nil, err
	}

	return &StoreConfig{
		DB: &DBConfig{
			Host: m[3],
			User: m[1],
			Pass: m[2],
			Name: m[5],
			Port: port,
		},
	}, nil
}

func (s *Shopware6) Version(docroot string) (string, error) {
	return getVersionFromComposer(docroot, sw6ComposerRgx)
}

func (s *Shopware6) BaseURLs(ctx context.Context, docroot string) ([]string, error) {
	cfg, err := s.ParseConfig(filepath.Join(docroot, s.ConfigPath()))
	if err != nil {
		return nil, err
	}

	db, err := ConnectDB(ctx, *cfg.DB)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT url FROM sales_channel_domain WHERE url NOT LIKE '%headless%'`)
	if err != nil {
		return nil, err
	}

	urls := []string{}
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err == nil {
			urls = append(urls, url)
		}
	}

	if len(urls) > 0 {
		return urls, nil
	}

	return nil, errors.New("base url(s) not found in database")
}
