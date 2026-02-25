package gocommerce

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type Magento1 struct {
	basePlatform
	Magerun string
}

const (
	versionRegex = "(?m)public static function getVersionInfo[^=]+=>\\s'(\\d)',[^=]+=>\\s'(\\d)',[^=]+=>\\s'(\\d)',[^=]+=>\\s'(\\d)',"
)

type m1Config struct {
	XMLName xml.Name       `xml:"config"`
	Global  m1ConfigGlobal `xml:"global"`
	Admin   m1ConfigAdmin  `xml:"admin"`
}

type m1ConfigGlobal struct {
	Resources m1ConfigResources `xml:"resources"`
}

type m1ConfigResources struct {
	DB           m1ConfigDB           `xml:"db"`
	DefaultSetup m1ConfigDefaultSetup `xml:"default_setup"`
}

type m1ConfigDB struct {
	TablePrefix string `xml:"table_prefix"`
}

type m1ConfigDefaultSetup struct {
	Connection m1ConfigConnection `xml:"connection"`
}

type m1ConfigConnection struct {
	Host     string `xml:"host"`
	Username string `xml:"username"`
	Password string `xml:"password"`
	DBName   string `xml:"dbname"`
}

type m1ConfigAdmin struct {
	Routers m1ConfigRouters `xml:"routers"`
}

type m1ConfigRouters struct {
	Adminhtml m1ConfigAdminhtml `xml:"adminhtml"`
}

type m1ConfigAdminhtml struct {
	Args m1ConfigArgs `xml:"args"`
}

type m1ConfigArgs struct {
	FrontName string `xml:"frontName"`
}

func (m1 *Magento1) ParseConfig(cfgPath string) (*StoreConfig, error) {
	xmlFile, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	var cfg m1Config
	if err := xml.NewDecoder(xmlFile).Decode(&cfg); err != nil {
		return nil, err
	}

	conn := cfg.Global.Resources.DefaultSetup.Connection
	if conn.Username == "" || conn.DBName == "" {
		return nil, fmt.Errorf("XML parse error for %s", cfgPath)
	}

	host := conn.Host
	port := 3306
	if h, p, e := parseHostPort(host); p > 0 && e == nil {
		port = p
		host = h
	}

	return &StoreConfig{
		DB: &DBConfig{
			Host:   host,
			User:   conn.Username,
			Pass:   conn.Password,
			Name:   conn.DBName,
			Prefix: cfg.Global.Resources.DB.TablePrefix,
			Port:   port,
		},
		AdminSlug: cfg.Admin.Routers.Adminhtml.Args.FrontName,
	}, nil
}

func (m1 *Magento1) BaseURLs(ctx context.Context, docroot string) ([]string, error) {
	cfgPath := filepath.Join(docroot, m1.ConfigPath())

	cfg, err := m1.ParseConfig(cfgPath)
	if err != nil {
		return nil, err
	}

	db, err := ConnectDB(ctx, *cfg.DB)
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
