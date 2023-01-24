package gocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindStoreAtRoot(t *testing.T) {
	root := fixtureBase + "/magento1"
	s := FindStoreAtRoot(root)
	assert.NotNil(t, s)
	assert.Equal(t, "Magento 1", s.Platform.Name())
}
