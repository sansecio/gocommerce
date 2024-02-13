package gocommerce

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/sansecio/gocommerce/phpcfg"
)

type (
	Magento2 struct {
		basePlatform
		Magerun string
	}

	composerRoot struct {
		Name    string            `json:"name"`
		Version string            `json:"version"`
		Require map[string]string `json:"require"`
	}

	composerPackages struct {
		Packages []struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"packages"`
	}
)

func (m2 *Magento2) ParseConfig(cfgPath string) (*StoreConfig, error) {
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
	if h, p, e := parseHostPort(host); p > 0 && e == nil {
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

func (m2 *Magento2) BaseURLs(docroot string) ([]string, error) {
	cfgPath := filepath.Join(docroot, m2.ConfigPath())

	urls, err := m2.getBaseURLsFromDatabase(cfgPath)
	if err == nil {
		return urls, nil
	}

	return m2.getBaseURLsFromConfig(cfgPath)
}

func (m2 *Magento2) Version(docroot string) (string, error) {
	version, err := getVersionFromLockFile(docroot + "/composer.lock")
	if err == nil {
		return version, nil
	}

	return getVersionFromJsonFile(docroot + "/composer.json")
}

func (m2 *Magento2) getBaseURLsFromConfig(cfgPath string) ([]string, error) {
	cm, err := phpcfg.ParsePath(cfgPath)
	if err != nil {
		return nil, err
	}

	r, err := regexp.Compile(`root.system.\w+.web.secure.base_url`)
	if err != nil {
		return nil, err
	}

	urls := []string{}
	for k, v := range cm {
		if r.MatchString(k) {
			urls = append(urls, v)
		}
	}

	if len(urls) > 0 {
		return urls, nil
	}

	return nil, errors.New("base url(s) not found in config")
}

func (m2 *Magento2) getBaseURLsFromDatabase(cfgPath string) ([]string, error) {
	cfg, err := m2.ParseConfig(cfgPath)
	if err != nil {
		return nil, err
	}

	db, err := ConnectDB(*cfg.DB)
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

func getVersionFromJsonFile(jsonFile string) (string, error) {
	jf, err := os.ReadFile(jsonFile)
	if err != nil {
		return "", err
	}

	cr := composerRoot{}
	err = json.Unmarshal(jf, &cr)
	if err != nil {
		return "", err
	}

	// First check if we can find a require on a system package.
	for p, v := range cr.Require {
		if isSystemPackage(p) {
			return v, nil
		}
	}

	// Then check if we are a git clone.
	if isRootPackage(cr.Name) {
		return cr.Version, nil
	}

	return "", errors.New("unable to determine version from composer.json")
}

func getVersionFromLockFile(lockFile string) (string, error) {
	lf, err := os.ReadFile(lockFile)
	if err != nil {
		return "", err
	}

	cp := composerPackages{}
	err = json.Unmarshal(lf, &cp)
	if err != nil {
		return "", err
	}

	for _, p := range cp.Packages {
		if isSystemPackage(p.Name) {
			return p.Version, nil
		}
	}

	return "", errors.New("unable to determine version from composer.lock")
}

func isRootPackage(packageName string) bool {
	match, _ := regexp.MatchString(`magento\/magento2...?`, packageName)
	return match
}

func isSystemPackage(packageName string) bool {
	match, _ := regexp.MatchString(`magento\/product-.*?-edition`, packageName)
	return match
}
