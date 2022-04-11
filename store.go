package gocommerce

import (
	"os"
	"path/filepath"
	"strings"
)

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

func docrootToStore(docroot string, pl PlatformInterface) *Store {
	cfgPath := filepath.Join(docroot, pl.ConfigPath())
	cfg, e := pl.ParseConfig(cfgPath)
	if e != nil {
		return nil
	}
	return &Store{docroot, pl, cfg}
}

func pathExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
