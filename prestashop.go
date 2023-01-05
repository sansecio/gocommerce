package gocommerce

import "github.com/sansecio/gocommerce/phpcfg"

func (p *Prestashop) ParseConfig(cfgPath string) (*StoreConfig, error) {
	cm, err := phpcfg.ParsePath(cfgPath)
	if err != nil {
		return nil, err
	}

	port := 3306 // root.parameters.database_port

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
