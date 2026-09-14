package gocommerce

import (
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// Servers with require_secure_transport=ON reject unencrypted connections with
// error 3159, which is what makes ConnectDB retry over TLS.
func TestSecureTransportRequired(t *testing.T) {
	err := &mysql.MySQLError{Number: mysqlerr.ER_SECURE_TRANSPORT_REQUIRED}
	require.True(t, secureTransportRequired(fmt.Errorf("ping failed: %w", err)))
	require.False(t, secureTransportRequired(&mysql.MySQLError{Number: 1045})) // access denied
	require.False(t, secureTransportRequired(errors.New("dial tcp: connection refused")))
}

func TestConnectDBDisabledDoesNotDial(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer listener.Close()
	accepted := make(chan error, 1)
	go func() {
		conn, err := listener.Accept()
		if err == nil {
			conn.Close()
		}
		accepted <- err
	}()

	addr := listener.Addr().(*net.TCPAddr)
	db, err := ConnectDB(nil, DBConfig{
		Host: "127.0.0.1",
		Port: addr.Port,
		User: "user",
		Pass: "password",
		Name: "shop",
	}, Options{SkipDatabase: true})
	require.ErrorIs(t, err, ErrDatabaseDisabled)
	require.Nil(t, db)
	require.NoError(t, listener.Close())
	require.Error(t, <-accepted)
}
