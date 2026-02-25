package gocommerce

import (
	"context"
	"errors"
	"time"
)

func prestashopBaseURLs(ctx context.Context, cfg *StoreConfig) ([]string, error) {
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

func prestashopVersion(cfg *StoreConfig) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := ConnectDB(ctx, *cfg.DB)
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
