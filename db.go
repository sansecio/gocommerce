package gocommerce

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" //nolint
)

var defaultSockets = []string{
	"/var/run/mysqld/mysqld.sock",
	"/var/lib/mysql/mysql.sock",
}

// NB copy StoreConfig, as we may modify it
func ConnectDB(ctx context.Context, cfg DBConfig) (*sql.DB, error) {
	// Mimic libmysql behavior, where "localhost" is overridden with
	// system specific unix socket.
	if cfg.Host == "localhost" || cfg.Host == "" {
		for _, s := range defaultSockets {
			if isSocket(s) {
				cfg.Host = s
				break
			}
		}
	}

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		return nil, err
	}

	if ctx == nil {
		ctx = context.Background()
	}
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func isSocket(path string) bool {
	s, e := os.Stat(path)
	if e != nil {
		return false
	}
	return s.Mode()&os.ModeSocket == os.ModeSocket
}
