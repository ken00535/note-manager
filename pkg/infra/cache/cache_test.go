package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	c *Cache
)

func TestMain(t *testing.M) {
	c = NewCache()
	t.Run()
}

func TestCache_Get_None(t *testing.T) {
	var val string
	err := c.Key("none").Get(&val)
	assert.Error(t, err, "")
}

func TestCache_Get_Struct(t *testing.T) {
	var val string
	type TC struct {
		Name string
		Num  float64
	}
	tc := TC{}
	tc.Num = 10
	tc.Name = "ken"
	c.Key("ken").Set(tc)
	err := c.Key("ken").Get(&val)
	assert.NoError(t, err, "")
	assert.Equal(t, "ken", val)
}
