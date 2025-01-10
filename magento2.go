package gocommerce

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/sansecio/gocommerce/phpcfg"
)

type (
	Magento2 struct {
		basePlatform
		Magerun string
	}
)

var m2ComposerRgx = regexp.MustCompile(`magento\/product-.*?-edition`)

func (m2 *Magento2) ParseConfig(cfgPath string) (*StoreConfig, error) {
	cm, err := phpcfg.ParsePath(cfgPath)
	if err != nil {
		return nil, err
	}

	port := 3306
	host := cm["root.db.connection.default.host"]

	if p, err := strconv.Atoi(cm["root.db.connection.default.port"]); err == nil {
		port = p
	}

	// Apparently you can stuff a port in the hostfield
	if h, p, e := parseHostPort(host); p > 0 && e == nil {
		port = p
		host = h
	}
	if cm["root.db.connection.default.username"] == "" || cm["root.db.connection.default.dbname"] == "" {
		return nil, fmt.Errorf("could not parse %s, missing user or db name", cfgPath)
	}

	return &StoreConfig{
		&DBConfig{
			Host:   host,
			User:   cm["root.db.connection.default.username"],
			Pass:   cm["root.db.connection.default.password"],
			Name:   cm["root.db.connection.default.dbname"],
			Prefix: cm["root.db.table_prefix"],
			Port:   port,
		},
		cm["root.backend.frontName"],
	}, nil
}

func (m2 *Magento2) BaseURLs(ctx context.Context, docroot string) ([]string, error) {
	cfgPath := filepath.Join(docroot, m2.ConfigPath())
	urls := []string{}
	if ud, err := m2.getBaseURLsFromDatabase(ctx, cfgPath); err == nil {
		urls = append(urls, ud...)
	}
	if uc, err := m2.getBaseURLsFromConfig(cfgPath); err == nil {
		urls = append(urls, uc...)
	}
	slices.Sort(urls)
	return slices.Compact(urls), nil
}

func (m2 *Magento2) Version(docroot string) (string, error) {
	return getVersionFromComposer(docroot, m2ComposerRgx)
}

func urlIsPlaceholder(url string) bool {
	return strings.HasPrefix(url, "{{") || strings.HasSuffix(url, "}}")
}

func (m2 *Magento2) getBaseURLsFromConfig(cfgPath string) ([]string, error) {
	cm, err := phpcfg.ParsePath(cfgPath)
	if err != nil {
		return nil, err
	}

	r, err := regexp.Compile(`root.system.\w+.web.(un)?secure.base_url`)
	if err != nil {
		return nil, err
	}

	urls := []string{}
	for k, v := range cm {
		if r.MatchString(k) && !urlIsPlaceholder(v) {
			urls = append(urls, v)
		}
	}

	if len(urls) > 0 {
		return urls, nil
	}

	return nil, errors.New("base url(s) not found in config")
}

func (m2 *Magento2) getBaseURLsFromDatabase(ctx context.Context, cfgPath string) ([]string, error) {
	cfg, err := m2.ParseConfig(cfgPath)
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
	rows, err := db.Query(`select distinct value from ` + prefix + `core_config_data where path like 'web/%secure/base_url'`)
	if err != nil {
		return nil, err
	}

	urls := []string{}
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err == nil && !urlIsPlaceholder(url) {
			urls = append(urls, url)
		}
	}

	if len(urls) > 0 {
		return urls, nil
	}

	return nil, errors.New("base url(s) not found in database")
}
