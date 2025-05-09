package gocommerce

import (
	"os"
	"path/filepath"
	"slices"
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
	"/httpdocs",  // youwe
	"/app/$USER", // adobe commerce
	"$HOME/var/$SITE_NAME/logs",
	"/var/www/html",
	"$HOME/html",                       // jetrails
	"/home/jetrails/*/html",            // jetrails
	"$HOME/applications/*/public_html", // cloudways
	"/data/web/public",                 // hypernode
	"/home/cloudpanel/htdocs/*/current",
	"/home/cloudpanel/htdocs/*",
	"/code",
	"/domains/*/http", // sonassi
	"/srv/*",
	"$HOME/domains/*/public_html",
	"/var/www/*",
	"/var/www/*/current",
	"/var/www/*/public_html",
	"/var/www/vhosts/*/htdocs",
	"/vhosts/*/httpdocs",   // plesk
	"$HOME/public_html/..", // nexcess
	"$HOME/public/..",      // hypernode
	"$HOME/public",         // hypernode, nimbus hosting
	"$HOME/??*.*",          // generic domain pattern
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

func findDocRoots() []string {
	var err error
	var roots []string
	for _, dr := range commonDocRoots {
		dr = os.ExpandEnv(dr)
		// Following is required to support "$HOME/public_html/.." in possible store paths.
		// There, public_html being a symlink to magento2/pub (Nexcess and others).
		// Without this lookup, public_html/.. would get lexically parsed by filepath.Clean()
		// ending up at the parent of public_html instead of magento2. --wdg
		if strings.Contains(dr, "/..") && !strings.Contains(dr, "*") && !strings.Contains(dr, "?") {
			dr, err = filepath.EvalSymlinks(dr)
			if err != nil { // eg some path component does not exist
				continue
			}
		}

		// extrapolate possible globs
		allPaths, err := filepath.Glob(dr)
		if err != nil {
			continue
		}
		roots = append(roots, allPaths...)
	}
	slices.Sort(roots)
	return slices.Compact(roots)
}

// DiscoverStores searches several common docroot locations for stores
func DiscoverStores() []*Store {
	stores := []*Store{}
	for _, p := range findDocRoots() {
		if s := FindStoreAtRoot(p); s != nil {
			stores = append(stores, s)
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
