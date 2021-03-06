package gocommerce

import (
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

	magento1 struct {
		basePlatform
		magerun string
	}

	magento2 struct {
		basePlatform
		Magerun string
	}
	// Wordpress struct {
	// 	baseSystem
	// }

	PlatformInterface interface {
		Name() string
		ParseConfig(cfgPath string) (*StoreConfig, error)
		// Version(docroot string) (string, error)
		// BaseURLs(docroot string) ([]string, error)
		ConfigPath() string
		UniquePath() string
	}
)

// func (b *basePlatform) Detect(docroot string) bool {
// 	return pathExists(filepath.Join(docroot, b.configPath))
// }
func (b *basePlatform) Name() string {
	return b.name
}
func (b *basePlatform) ConfigPath() string {
	return b.configPath
}
func (b *basePlatform) UniquePath() string {
	return b.uniquePath
}

var (
	Magento1 = magento1{
		basePlatform{
			"Magento1",
			"app/etc/local.xml",
			"app/etc/local.xml",
		},
		"n98-magerun"}

	Magento2 = magento2{
		basePlatform{
			"Magento2",
			"app/etc/env.php",
			"app/etc/env.php",
		},
		"n98-magerun2"}

	AllPlatforms = []PlatformInterface{
		&Magento1,
		&Magento2,
	} // {"Magento1", "/app/etc/local.xml", "n98-magerun", ParseMagento1Config},
	// {"Magento2", "/app/etc/env.php", "n98-magerun2", nil},

)

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
