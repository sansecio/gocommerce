package gocommerce

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/xmlpath.v2"
)

type Magento1 struct {
	basePlatform
	Magerun string
}

const (
	versionRegex = "(?m)public static function getVersionInfo[^=]+=>\\s'(\\d)',[^=]+=>\\s'(\\d)',[^=]+=>\\s'(\\d)',[^=]+=>\\s'(\\d)',"
)

func (m1 *Magento1) ParseConfig(cfgPath string) (*StoreConfig, error) {
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
	user, _ := xmlpath.MustCompile("/config/global/resources/default_setup/connection/username").String(root)
	pass, _ := xmlpath.MustCompile("/config/global/resources/default_setup/connection/password").String(root)
	host, _ := xmlpath.MustCompile("/config/global/resources/default_setup/connection/host").String(root)
	dbname, _ := xmlpath.MustCompile("/config/global/resources/default_setup/connection/dbname").String(root)
	slug, _ := xmlpath.MustCompile("/config/admin/routers/adminhtml/args/frontName").String(root)

	if user == "" || dbname == "" {
		return nil, fmt.Errorf("XML parse error for %s", cfgPath)
	}

	port := 3306
	if h, p, e := parseHostPort(host); p > 0 && e == nil {
		port = p
		host = h
	}

	return &StoreConfig{
		DB: &DBConfig{
			Host:   host,
			User:   user,
			Pass:   pass,
			Name:   dbname,
			Prefix: prefix,
			Port:   port,
		},
		AdminSlug: slug,
	}, nil
}

func (m1 *Magento1) BaseURLs(docroot string, ctx context.Context) ([]string, error) {
	cfgPath := filepath.Join(docroot, m1.ConfigPath())

	cfg, err := m1.ParseConfig(cfgPath)
	if err != nil {
		return nil, err
	}

	db, err := ConnectDB(*cfg.DB, ctx)
	if err != nil {
		return nil, err
	}

	prefix, err := cfg.DB.SafePrefix()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(`select distinct value from ` + prefix + `core_config_data where path like 'web/%secure/base_url'`)
	if err != nil {
		return nil, err
	}

	urls := []string{}
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err == nil {
			urls = append(urls, url)
		}
	}

	if len(urls) > 0 {
		return urls, nil
	}

	return nil, errors.New("base url(s) not found in database")
}

func (m1 *Magento1) Version(docroot string) (string, error) {
	dat, err := os.ReadFile(filepath.Join(docroot, "app/Mage.php"))
	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(versionRegex)
	m := r.FindAllStringSubmatch(string(dat), -1)
	if len(m) == 0 || len(m[0]) < 5 {
		return "", errors.New("could not determine magento 1 version")
	}

	return fmt.Sprintf("%s.%s.%s.%s", m[0][1], m[0][2], m[0][3], m[0][4]), nil
}
