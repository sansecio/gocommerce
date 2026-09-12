package gocommerce

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// Servers with require_secure_transport=ON reject unencrypted connections with
// error 3159, which is what makes ConnectDB retry over TLS.
func TestSecureTransportRequired(t *testing.T) {
	err := &mysql.MySQLError{Number: erSecureTransportRequired}
	require.True(t, secureTransportRequired(fmt.Errorf("ping failed: %w", err)))
	require.False(t, secureTransportRequired(&mysql.MySQLError{Number: 1045})) // access denied
	require.False(t, secureTransportRequired(errors.New("dial tcp: connection refused")))
}
