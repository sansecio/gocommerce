package gocommerce

import (
	"context"
	"database/sql"
	"errors"
	"os"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
)

var ErrDatabaseDisabled = errors.New("database access disabled")

var defaultSockets = []string{
	"/var/run/mysqld/mysqld.sock",
	"/var/lib/mysql/mysql.sock",
}

// NB copy StoreConfig, as we may modify it
func ConnectDB(ctx context.Context, cfg DBConfig, opts Options) (*sql.DB, error) {
	if opts.SkipDatabase {
		return nil, ErrDatabaseDisabled
	}

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

	if ctx == nil {
		ctx = context.Background()
	}

	db, err := dial(ctx, cfg.DSN())
	if err != nil && secureTransportRequired(err) {
		// retry encrypted
		db, err = dial(ctx, cfg.DSN()+"&tls=preferred")
	}
	return db, err
}

func dial(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}

func secureTransportRequired(err error) bool {
	var e *mysql.MySQLError
	return errors.As(err, &e) && e.Number == mysqlerr.ER_SECURE_TRANSPORT_REQUIRED
}

func isSocket(path string) bool {
	s, e := os.Stat(path)
	if e != nil {
		return false
	}
	return s.Mode()&os.ModeSocket == os.ModeSocket
}
