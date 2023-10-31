package gocommerce

import (
	"os"
	"path/filepath"
	"strings"
)

var commonDocRoots = []string{
	"$PWD", // convenient when manually searching for current dir
	"$DOCROOT",
	"$HOME/public_html/deployments/current",
	"$HOME/public_html/current",
	"$HOME/public_html",
	"$HOME/www",
	"$HOME/httpdocs",
	"$HOME/master/Magento1Website",
	"$HOME/current",
	"$HOME/magento2",
	"/app/$USER", // adobe commerce
	"$HOME/var/$SITE_NAME/logs",
	"/var/www/html",
	"$HOME/html",                       // jetrails
	"/home/jetrails/*/html",            // jetrails
	"$HOME/applications/*/public_html", // cloudways
	"/data/web/public",                 // hypernode
	"/home/cloudpanel/htdocs/*/current",
	"/code",
	"/domains/*/http", // sonassi
	"/srv/*",
	"$HOME/domains/*/public_html",
}

func (s *Store) ConfigPath() string {
	return filepath.Join(s.DocRoot, s.Platform.ConfigPath())
}

func FindStoreAtRoot(docroot string) *Store {
	for _, pl := range AllPlatforms {
		path := filepath.Join(docroot, pl.UniquePath())
		if !pathExists(path) {
			continue
		}
		if s := docrootToStore(docroot, pl); s != nil {
			return s
		}
	}
	return nil
}

// For recursively walking the filesystem:
// derive a store from a uniquely identifying path
func FindStoreAtUniquePath(path string) *Store {
	for _, pl := range AllPlatforms {
		if !strings.HasSuffix(path, pl.UniquePath()) {
			continue
		}
		docroot := path[:len(path)-len(pl.UniquePath())-1]
		if s := docrootToStore(docroot, pl); s != nil {
			return s
		}

	}
	return nil
}

// DiscoverStores searches several common docroot locations for stores
func DiscoverStores() []*Store {
	stores := []*Store{}
	for _, dr := range commonDocRoots {
		dr = os.ExpandEnv(dr)
		// extrapolate possible globs
		allPaths, err := filepath.Glob(dr)
		if err != nil {
			continue
		}
		for _, p := range allPaths {
			if s := FindStoreAtRoot(p); s != nil {
				stores = append(stores, s)
			}
		}
	}
	return stores
}

func docrootToStore(docroot string, pl PlatformInterface) *Store {
	cfgPath := filepath.Join(docroot, pl.ConfigPath())
	cfg, _ := pl.ParseConfig(cfgPath)
	return &Store{docroot, pl, cfg}
}

func pathExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
