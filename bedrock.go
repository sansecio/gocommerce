package gocommerce

import (
	"os"
	"regexp"
)

type Bedrock struct {
	basePlatform
}

var brLookupRgx = map[string]string{
	"user":   `DB_USER=['"]?(\S+)['"]?`,
	"pass":   `DB_PASSWORD=['"]?(\S+)['"]?`,
	"host":   `DB_HOST=['"]?(\S+)['"]?`,
	"db":     `DB_NAME=['"]?(\S+)['"]?`,
	"prefix": `DB_PREFIX=['"]?(\S+)['"]?`,
}

func (w *Bedrock) ParseConfig(cfgPath string) (*StoreConfig, error) {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	matches := map[string]string{}
	port := 3306

	for k, v := range brLookupRgx {
		m := regexp.MustCompile(v).FindStringSubmatch(string(data))
		if len(m) != 2 {
			continue
		}
		matches[k] = m[1]
	}

	if h, p, e := parseHostPort(matches["host"]); p > 0 && e == nil {
		matches["host"] = h
		port = p
	}

	return &StoreConfig{
		DB: &DBConfig{
			Host:   matches["host"],
			User:   matches["user"],
			Pass:   matches["pass"],
			Name:   matches["db"],
			Prefix: matches["prefix"],
			Port:   port,
		},
	}, nil
}
