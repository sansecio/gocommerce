package gocommerce

import (
	"fmt"
	"testing"
)

const fixtureBase = "fixture/"

func dbConfigFromSource(t *testing.T, src string, pl PlatformInterface) *DBConfig {
	cfg, e := pl.ParseConfig(src)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	return cfg.DB
}
