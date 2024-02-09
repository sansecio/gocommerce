package gocommerce

import (
	"testing"
)

var woocommerce = WooCommerce{}

func TestWooCommerceConfigToDSN(t *testing.T) {
	tests := map[string]DBConfig{
		"wp-config.php": {
			Host:   "db.hostname.com",
			Port:   3306,
			Name:   "my_db_name",
			User:   "my_db_user",
			Pass:   "my_password",
			Prefix: "wp_shw5n4a6zq_",
		},
		"wp-config2.php": {
			Host:   "xxhost",
			Port:   3306,
			Name:   "xx_db",
			User:   "xxxuser",
			Pass:   "xxpass",
			Prefix: "wp102_",
		},
		"wp-config3.php": {
			Host:   "localhost",
			Port:   3306,
			Name:   "s87dfs8dfsdf",
			User:   "Ad87d8asdhads",
			Pass:   "Â skldjfksdfkjsdf",
			Prefix: "myorefx999_",
		},
		"wp-config4.php": {
			Host:   "10.10.22.203",
			Port:   3306,
			Name:   "db_cgq205qk4dw",
			User:   "user_cgq205qk4dw",
			Pass:   "202b2760-3673-4fd0-8b1b-8bcf2bd0c060",
			Prefix: "sdg_",
		},
		"wp-config5.php": {
			Host:   "127.0.0.1",
			Port:   3306,
			Name:   "live_db",
			User:   "live_user",
			Pass:   "xyz123",
			Prefix: "wp_",
		},
		"wp-config6.php": {
			Host:   "127.0.0.1",
			Port:   3308,
			Name:   "live_db",
			User:   "live_user",
			Pass:   "xyz123",
			Prefix: "wp_",
		},
	}

	for cnf, want := range tests {
		path := fixtureBase + "/wordpress/configs/" + cnf
		got := dbConfigFromSource(t, path, &woocommerce)

		if got.DSN() != want.DSN() {
			t.Errorf("%v: ConfigToDSN() = %v, want %v", cnf, got, want)
		}

		if got.Prefix != want.Prefix {
			t.Errorf("%v: Prefix do not match: <%v> vs <%v>", cnf, got.Prefix, want.Prefix)
		}

	}
}
