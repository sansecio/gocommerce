package gocommerce

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sansecio/gocommerce/phpcfg"
)

func (m2 *magento2) ParseConfig(cfgPath string) (*StoreConfig, error) {
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
	if h, p := hostToHostPort(host); p > 0 {
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

func hostToHostPort(host string) (string, int) {
	if !strings.Contains(host, ":") {
		return host, 0
	}
	t := strings.Split(host, ":")
	if len(t) != 2 {
		return t[0], 0
	}
	if p, err := strconv.Atoi(t[1]); err == nil {
		return t[0], p
	}
	return t[0], 0
}
