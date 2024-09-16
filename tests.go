package gocommerce

import (
	"fmt"
	"testing"
)

const fixtureBase = "fixture/"

func dbConfigFromSource(_ *testing.T, src string, pl PlatformInterface) *DBConfig {
	cfg, e := pl.ParseConfig(src)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	return cfg.DB
}
