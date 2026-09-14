package gocommerce

import (
	"context"
	"path/filepath"
	"regexp"
)

type Sylius struct {
	basePlatform
}

// anchored so it matches sylius/sylius but not sylius/sylius-rector et al.
var syliusComposerRgx = regexp.MustCompile(`^sylius/sylius$`)

func (s *Sylius) ParseConfig(cfgPath string, opts Options) (*StoreConfig, error) {
	return symfonyParseConfig(cfgPath, opts)
}

func (s *Sylius) BaseURLs(ctx context.Context, docroot string, opts Options) ([]string, error) {
	cfg, err := s.ParseConfig(filepath.Join(docroot, s.ConfigPath()), opts)
	if err != nil {
		return nil, err
	}

	// sylius_channel stores bare hostnames, not full URLs
	hosts, err := symfonyColumnURLs(ctx, cfg, `SELECT hostname FROM sylius_channel WHERE enabled = 1 AND hostname IS NOT NULL`, opts)
	if err != nil {
		return nil, err
	}

	for i, h := range hosts {
		hosts[i] = "https://" + h + "/"
	}
	return hosts, nil
}

func (s *Sylius) Version(docroot string, opts Options) (string, error) {
	return getVersionFromComposer(docroot, syliusComposerRgx)
}
