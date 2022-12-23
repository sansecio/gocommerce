package gocommerce

import "github.com/sansecio/gocommerce/phpcfg"

func (p *prestashop) TableChecks() []TableCheck {
	return []TableCheck{
		{"configuration", "name", "value", ""},
		{"cms_lang", "id_cms", "content", ""},
		{"employee", "'email'", "email", ""},
		{"vccontentanywhere_lang", "id_vccontentanywhere", "CONCAT_WS(' ',title,content)", ""},
	}
}

func (p *prestashop) ParseConfig(cfgPath string) (*StoreConfig, error) {
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
