package gocommerce

import (
	"errors"
	"fmt"
	"strings"
)

type (
	DBConfig struct {
		Host   string
		User   string
		Pass   string
		Name   string
		Prefix string
		Port   int
	}

	StoreConfig struct {
		DB        *DBConfig
		AdminSlug string
	}

	Store struct {
		DocRoot  string
		Platform PlatformInterface
		Config   *StoreConfig
	}

	basePlatform struct {
		name       string
		configPath string
		uniquePath string // A relative path that is sufficiently unique to identify a particular platform
	}

	Magento1 struct {
		basePlatform
		Magerun string
	}

	Magento2 struct {
		basePlatform
		Magerun string
	}

	Shopware5 struct {
		basePlatform
	}

	Shopware6 struct {
		basePlatform
	}

	Prestashop struct {
		basePlatform
	}

	Woocommerce struct {
		basePlatform
	}

	PlatformInterface interface {
		Name() string
		ParseConfig(cfgPath string) (*StoreConfig, error)
		Version(docroot string) (string, error)
		BaseURLs(docroot string) ([]string, error)
		ConfigPath() string
		UniquePath() string
	}
)

var (
	AllPlatforms = []PlatformInterface{
		&Magento1{
			basePlatform{
				"Magento1",
				"app/etc/local.xml",
				"app/etc/local.xml",
			},
			"n98-magerun",
		},
		&Magento2{
			basePlatform{
				"Magento2",
				"app/etc/env.php",
				"app/etc/env.php",
			},
			"n98-magerun2",
		},
		&Shopware5{
			basePlatform{
				"Shopware 5",
				"config.php",
				"engine/Shopware/Application.php",
			},
		},
		&Shopware6{
			basePlatform{
				"Shopware 6",
				".env",
				"vendor/shopware/core/Framework/ShopwareException.php",
			},
		},
		&Prestashop{
			basePlatform{
				"Prestashop",
				"app/config/parameters.php",
				"app/config/parameters.php",
			},
		},
		&Woocommerce{
			basePlatform{
				"Woocommerce",
				"wp-config.php",
				"wp-config.php",
			},
		},
	}
)

func (b *basePlatform) Name() string {
	return b.name
}

func (b *basePlatform) ConfigPath() string {
	return b.configPath
}

func (b *basePlatform) UniquePath() string {
	return b.uniquePath
}

func (b *basePlatform) ParseConfig(cfgPath string) (*StoreConfig, error) {
	return nil, errors.New("not implemented")
}

func (b *basePlatform) BaseURLs(docroot string) ([]string, error) {
	return nil, errors.New("not implemented")
}

func (b *basePlatform) Version(docroot string) (string, error) {
	return "", errors.New("not implemented")
}

func (c *DBConfig) DSN() string {
	if c.User == "" || c.Name == "" {
		return ""
	}
	host := c.Host
	if host == "" {
		host = "localhost"
	}
	port := c.Port
	if port == 0 {
		port = 3306
	}

	var network, address string
	if strings.Contains(c.Host, "/") {
		network = "unix"
		address = c.Host
	} else {
		network = "tcp"
		address = fmt.Sprintf("%s:%d", host, port)
	}

	// oldpasswords = Required for MySQL 4.0 servers, ugh...
	return fmt.Sprintf("%s:%s@%s(%s)/%s?allowOldPasswords=true",
		c.User,
		c.Pass,
		network,
		address,
		c.Name)
}
