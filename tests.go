package gocommerce

import (
	"fmt"
	"testing"
)

const fixtureBase = "fixture/"

// platformByName returns the fully-configured platform instance registered in
// AllPlatforms, so tests exercise the same configPath/uniquePath as production.
func platformByName(t *testing.T, name string) PlatformInterface {
	for _, pl := range AllPlatforms {
		if pl.Name() == name {
			return pl
		}
	}
	t.Fatalf("platform %q not registered in AllPlatforms", name)
	return nil
}

func dbConfigFromSource(_ *testing.T, src string, pl PlatformInterface) *DBConfig {
	cfg, e := pl.ParseConfig(src)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	return cfg.DB
}
