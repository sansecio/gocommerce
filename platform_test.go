package gocommerce

import (
	"testing"
)

func TestSafePrefix(t *testing.T) {
	dummyConfig := &DBConfig{
		Host: "localhost",
		Port: 3306,
		Name: "magento",
		User: "romeo",
		Pass: "juliet",
	}

	validPrefixes := []string{
		"",
		"mage2_",
	}
	invalidPrefixes := []string{
		"core_config_data; DROP TABLE core_config_data --",
	}

	for _, prefix := range validPrefixes {
		dummyConfig.Prefix = prefix
		if _, err := dummyConfig.SafePrefix(); err != nil {
			t.Errorf("expected prefix `%v` to be valid, but got error: `%v`", prefix, err)
		}
	}

	for _, prefix := range invalidPrefixes {
		dummyConfig.Prefix = prefix
		if _, err := dummyConfig.SafePrefix(); err == nil {
			t.Errorf("expected prefix `%v` to be invalid, but got no error", prefix)
		}
	}
}
