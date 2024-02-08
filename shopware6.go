package gocommerce

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Shopware6 struct {
	basePlatform
}

const (
	// DATABASE_URL=mysql://db-user-1:rhPb5xC2242444mFZDB@localhost:3306/db-1
	DBURL = `(?m)^\s*DATABASE_URL\s*="?\s*mysql://(.+?):(.+?)@(.+?):(\d+)/(.+?)"?$`
)

func (s *Shopware6) ParseConfig(cfgPath string) (*StoreConfig, error) {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	m := regexp.MustCompile(DBURL).FindStringSubmatch(string(data))
	if len(m) != 6 {
		return nil, fmt.Errorf("could not parse DSN in Shopware config path %s", cfgPath)
	}

	port, err := strconv.Atoi(m[4])
	if err != nil {
		return nil, err
	}

	return &StoreConfig{
		DB: &DBConfig{
			Host: m[3],
			User: m[1],
			Pass: m[2],
			Name: m[5],
			Port: port,
		},
	}, nil
}
