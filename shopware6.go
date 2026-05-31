package gocommerce

import (
	"context"
	"path/filepath"
	"regexp"
)

type Shopware6 struct {
	basePlatform
}

var sw6ComposerRgx = regexp.MustCompile(`shopware\/core`)

func (s *Shopware6) ParseConfig(cfgPath string) (*StoreConfig, error) {
	return symfonyParseConfig(cfgPath)
}

func (s *Shopware6) Version(docroot string) (string, error) {
	return getVersionFromComposer(docroot, sw6ComposerRgx)
}

func (s *Shopware6) BaseURLs(ctx context.Context, docroot string) ([]string, error) {
	cfg, err := s.ParseConfig(filepath.Join(docroot, s.ConfigPath()))
	if err != nil {
		return nil, err
	}
	return symfonyColumnURLs(ctx, cfg,
		`SELECT url FROM sales_channel_domain WHERE url NOT LIKE '%headless%'`)
}
