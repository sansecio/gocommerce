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
	XMLName xml.Name `xml:"config"`
	Global  struct {
		Resources struct {
			DB struct {
				TablePrefix string `xml:"table_prefix"`
			} `xml:"db"`
			DefaultSetup struct {
				Connection struct {
					Host     string `xml:"host"`
					Username string `xml:"username"`
					Password string `xml:"password"`
					DBName   string `xml:"dbname"`
				} `xml:"connection"`
			} `xml:"default_setup"`
		} `xml:"resources"`
	} `xml:"global"`
	Admin struct {
		Routers struct {
			Adminhtml struct {
				Args struct {
					FrontName string `xml:"frontName"`
				} `xml:"args"`
			} `xml:"adminhtml"`
		} `xml:"routers"`
	} `xml:"admin"`
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

	host, port := conn.Host, 3306
	if h, p, e := parseHostPort(conn.Host); p > 0 && e == nil {
		host, port = h, p
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
