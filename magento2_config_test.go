package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfigMultiDB(t *testing.T) {
	m2 := Magento2{}
	dbc := dbConfigFromSource(t, fixtureBase+"/magento2/configs/multidb.php", &m2)
	assert.Equal(t, "goodhost", dbc.Host)
	assert.Equal(t, "gooddb", dbc.Name)
	assert.Equal(t, "gooduser", dbc.User)
	assert.Equal(t, "goodpass", dbc.Pass)
	assert.Equal(t, 3306, dbc.Port)
	assert.Equal(t, "", dbc.Prefix)
}

func TestParseConfigSimpleDB(t *testing.T) {
	m2 := Magento2{}
	dbc := dbConfigFromSource(t, fixtureBase+"/magento2/configs/simple.php", &m2)
	assert.Equal(t, "goodhost", dbc.Host)
	assert.Equal(t, "gooddb", dbc.Name)
	assert.Equal(t, "gooduser", dbc.User)
	assert.Equal(t, "verylongpassword", dbc.Pass)
	assert.Equal(t, 3306, dbc.Port)
	assert.Equal(t, "pref_", dbc.Prefix)
	assert.Equal(t, "gooduser:verylongpassword@tcp(goodhost:3306)/gooddb?allowOldPasswords=true", dbc.DSN())
}

func TestParseConfigEmpty(t *testing.T) {
	m2 := Magento2{}
	dbc := dbConfigFromSource(t, fixtureBase+"/magento2/configs/empty.php", &m2)
	assert.Nil(t, dbc)
}

func TestActualConfig(t *testing.T) {
	m2 := Magento2{}
	dbc := dbConfigFromSource(t, fixtureBase+"/magento2/app/etc/env.php", &m2)
	assert.Equal(t, "", dbc.Prefix)
	assert.Equal(t, "app:sldfjlskdfklds@tcp(localhost:3306)/magento2?allowOldPasswords=true", dbc.DSN())
}

func TestCrash1(t *testing.T) {
	m2 := Magento2{}
	dbc := dbConfigFromSource(t, fixtureBase+"/magento2/configs/crash1.php", &m2)
	assert.Equal(t, "xx:xx@tcp(localhost:3306)/xx?allowOldPasswords=true", dbc.DSN())
}

func TestCrash2(t *testing.T) {
	m2 := Magento2{}
	dbc := dbConfigFromSource(t, fixtureBase+"/magento2/configs/crash2.php", &m2)
	assert.Equal(t, "xx:xx@tcp(10.10.20.39:3306)/xx?allowOldPasswords=true", dbc.DSN())
}

func TestPHPParserDoesNotChokeOnListItems(t *testing.T) {
	m2 := Magento2{}
	dbc := dbConfigFromSource(t, fixtureBase+"/magento2/configs/crash3.php", &m2)
	assert.Equal(t, "myuser:mypass@tcp(myhost:3306)/mydb?allowOldPasswords=true", dbc.DSN())
}

func TestPortInHost(t *testing.T) {
	m2 := Magento2{}
	dbc := dbConfigFromSource(t, fixtureBase+"/magento2/configs/hostport.php", &m2)
	assert.Equal(t, "gooduser:verylongpassword@tcp(goodhost:3309)/gooddb?allowOldPasswords=true", dbc.DSN())
}

func TestScanNonExistantFile(t *testing.T) {
	m2 := Magento2{}
	dbc := dbConfigFromSource(t, "/do/not/exist", &m2)
	assert.Nil(t, dbc)
}
