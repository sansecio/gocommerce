package gocommerce

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/xmlpath.v2"
)

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

func (m1 *Magento1) BaseURLs(docroot string) ([]string, error) {
	cfgPath := filepath.Join(docroot, m1.ConfigPath())

	cfg, err := m1.ParseConfig(cfgPath)
	if err != nil {
		return nil, err
	}

	db, err := ConnectDB(*cfg.DB)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`select distinct value from core_config_data where path like 'web/%secure/base_url'`)
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
