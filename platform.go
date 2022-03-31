package gocommerce

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type (
	ConfigParser func(path string) (*StoreConfig, error)
	StoreConfig  struct {
		DBHost    string
		DBUser    string
		DBPass    string
		DBName    string
		DBPrefix  string
		DBPort    int
		AdminSlug string
	}
	Store struct {
		Platform *Platform
		DocRoot  string
		Config   *StoreConfig
	}

	Platform struct {
		Name         string
		ConfigFile   string
		Magerun      string
		ConfigParser ConfigParser
	}
)

var (
	AllPlatforms = []*Platform{
		{"Magento1", "/app/etc/local.xml", "n98-magerun", ParseMagento1Config},
		{"Magento2", "/app/etc/env.php", "n98-magerun2", nil},
	}
)

func (s *Store) ConfigPath() string {
	return filepath.Join(s.DocRoot, s.Platform.ConfigFile)
}

func (cfg *StoreConfig) DSN() string {
	if cfg.DBUser == "" || cfg.DBName == "" {
		return ""
	}
	host := cfg.DBHost
	if host == "" {
		host = "localhost"
	}
	port := cfg.DBPort
	if port == 0 {
		port = 3306
	}

	var network, address string
	if strings.Contains(cfg.DBHost, "/") {
		network = "unix"
		address = cfg.DBHost
	} else {
		network = "tcp"
		address = fmt.Sprintf("%s:%d", host, port)
	}

	// oldpasswords = Required for MySQL 4.0 servers, ugh...
	return fmt.Sprintf("%s:%s@%s(%s)/%s?allowOldPasswords=true",
		cfg.DBUser,
		cfg.DBPass,
		network,
		address,
		cfg.DBName)
}

func FindStore(docroot string) *Store {
	for _, pl := range AllPlatforms {
		path := filepath.Join(docroot, pl.ConfigFile)
		if !pathExists(path) {
			continue
		}
		var cfg *StoreConfig
		if pl.ConfigParser != nil {
			cfg, _ = pl.ConfigParser(path)
			if cfg == nil { // parse error
				continue
			}
		}
		return &Store{pl, docroot, cfg}
	}
	return nil
}

func pathExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
