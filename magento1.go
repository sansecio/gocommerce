package gocommerce

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/xmlpath.v2"
)

func ParseMagento1Config(path string) (*StoreConfig, error) {
	xmlFile, err := os.Open(path)
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
		return nil, fmt.Errorf("XML parse error for %s", path)
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
		DBHost:    h,
		DBUser:    u,
		DBPass:    p,
		DBName:    dbname,
		DBPrefix:  prefix,
		DBPort:    port,
		AdminSlug: slug,
	}, nil
}
