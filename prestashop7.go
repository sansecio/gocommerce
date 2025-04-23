package gocommerce

import (
	"strconv"

	"github.com/sansecio/gocommerce/phpcfg"
)

type Prestashop7 struct {
	basePlatform
}

func (p *Prestashop7) ParseConfig(cfgPath string) (*StoreConfig, error) {
	cm, err := phpcfg.ParsePath(cfgPath)
	if err != nil {
		return nil, err
	}

	port := 3306
	if dbPort := cm["root.parameters.database_port"]; dbPort != "" {
		if pi, err := strconv.Atoi(dbPort); err == nil {
			port = pi
		}
	}

	return &StoreConfig{
		DB: &DBConfig{
			Host:   cm["root.parameters.database_host"],
			User:   cm["root.parameters.database_user"],
			Pass:   cm["root.parameters.database_password"],
			Name:   cm["root.parameters.database_name"],
			Prefix: cm["root.parameters.database_prefix"],
			Port:   port,
		},
	}, nil
}
