package gocommerce

import "errors"

func (m1 *prestashop) ParseConfig(cfgPath string) (*StoreConfig, error) {
	return nil, errors.New("not implemented") // TODO
}

func (m1 *prestashop) BaseURLs(docroot string) ([]string, error) {
	return nil, errors.New("not implemented") // TODO
}

func (m1 *prestashop) Version(docroot string) (string, error) {
	return "", errors.New("not implemented") // TODO
}
