package gocommerce

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/xmlpath.v2"
)

// func (m1 *magento1) BaseURLs(docroot string) ([]string, error) {
// 	return []string{}, fmt.Errorf("not implemented yet")
// }

// func (m1 *magento1) Version(docroot string) (string, error) {
// 	return "", fmt.Errorf("not implemented yet")
// }

func (m1 *magento1) ParseConfig(cfgPath string) (*StoreConfig, error) {
	xmlFile, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	root, err := xmlpath.Parse(xmlFile)
	if err != nil {
		return nil, err
	}
	// path := xmlpath.MustCompile("/config/global/resources/db/table_prefix")

	prefix, _ := xmlpath.MustCompile("/config/global/resources/db/table_prefix").String(root)
	u, _ := xmlpath.MustCompile("/config/global/resources/default_setup/connection/username").String(root)
	p, _ := xmlpath.MustCompile("/config/global/resources/default_setup/connection/password").String(root)
	h, _ := xmlpath.MustCompile("/config/global/resources/default_setup/connection/host").String(root)
	dbname, _ := xmlpath.MustCompile("/config/global/resources/default_setup/connection/dbname").String(root)
	slug, _ := xmlpath.MustCompile("/config/admin/routers/adminhtml/args/frontName").String(root)
	port := 3306

	if u == "" || dbname == "" {
		return nil, fmt.Errorf("XML parse error for %s", cfgPath)
	}

	// if strings.Contains(h, ":") {
	token := strings.SplitN(h, ":", 2)
	if len(token) == 2 {
		h = token[0]
		port, err = strconv.Atoi(token[1])
		if err != nil || port <= 0 || port > 65536 {
			port = 3306
		}
	}
	// }

	return &StoreConfig{
		DB: &DBConfig{
			Host:   h,
			User:   u,
			Pass:   p,
			Name:   dbname,
			Prefix: prefix,
			Port:   port,
		},
		AdminSlug: slug,
	}, nil
}
