package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfigSimpleMagento1DB(t *testing.T) {
	m1 := Magento1{}
	dbc := dbConfigFromSource(t, fixtureBase+"/magento1/app/etc/local.xml", &m1)

	assert.Equal(t, "localhost", dbc.Host)
	assert.Equal(t, "mag1", dbc.Name)
	assert.Equal(t, "app", dbc.User)
	assert.Equal(t, "sdkjfhjksdfjk", dbc.Pass)
	assert.Equal(t, 3306, dbc.Port)
	assert.Equal(t, "", dbc.Prefix)
}

func TestParseConfigHostWithPort(t *testing.T) {
	m1 := Magento1{}
	dbc := dbConfigFromSource(t, fixtureBase+"/magento1/configs/app/etc/local.xml.hostport", &m1)
	assert.Equal(t, "localhost", dbc.Host)
	assert.Equal(t, 3307, dbc.Port)
}

func TestParseEmptyPassword(t *testing.T) {
	m1 := Magento1{}
	dbc := dbConfigFromSource(t, fixtureBase+"/magento1/configs/app/etc/local.xml.nopass", &m1)
	assert.Equal(t, "jeroen:@tcp(db:3306)/jeroen_schweigmann?allowOldPasswords=true", dbc.DSN())
}

// func TestGoogleCloudSocket(t *testing.T) {
// TODO: Why did we fail colons in host?
// dbc := dbConfigFromSource(t, fixtureBase+"/magento1/configs/app/etc/local.xml.google-cloud-socket", &magento1)
// assert.Equal(t, "app:M7PgT65EbGyIooZsxyV6Az7TgXDS5F@unix(/cloudsql/mystore:europe-west1:magento)/mag1?allowOldPasswords=true", dbc.DSN())
// }
