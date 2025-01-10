package gocommerce

import (
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"strings"
)

type (
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

func getVersionFromComposerJSONFile(
	jsonFile string,
	pkgRgx *regexp.Regexp,
) (string, error) {
	jf, err := os.ReadFile(jsonFile)
	if err != nil {
		return "", err
	}

	cr := composerRoot{}
	err = json.Unmarshal(jf, &cr)
	if err != nil {
		return "", err
	}

	for p, v := range cr.Require {
		if pkgRgx.MatchString(p) {
			return v, nil
		}
	}

	return "", errors.New("unable to determine version from composer.json")
}

func getVersionFromComposerLockFile(
	lockFile string,
	pkgRgx *regexp.Regexp,
) (string, error) {
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
		if pkgRgx.MatchString(p.Name) {
			return p.Version, nil
		}
	}

	return "", errors.New("unable to determine version from composer.lock")
}

func stripVersionPrefix(v string) string {
	if strings.HasPrefix(v, "v") {
		return v[1:]
	}
	return v
}

func getVersionFromComposer(
	docroot string,
	pkgRgx *regexp.Regexp,
) (string, error) {
	version, err := getVersionFromComposerLockFile(docroot+"/composer.lock", pkgRgx)
	if err == nil {
		return stripVersionPrefix(version), nil
	}
	version, err = getVersionFromComposerJSONFile(docroot+"/composer.json", pkgRgx)
	if err == nil {
		return stripVersionPrefix(version), nil
	}
	return "", err
}
