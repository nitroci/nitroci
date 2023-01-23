package configs

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestEnsureConfigurationFailIfNotExist(t *testing.T) {
  assert := assert.New(t)

  cfgfile, err := EnsureConfiguration("./.nitroci/config.ini")

  assert.NotNil(err)
  assert.Empty(cfgfile)
}

func Add(x, y int) (res int) {
	return x + y
}

func BenchmarkAdd(b *testing.B){
    for i :=0; i < b.N ; i++{
        Add(4, 6)
    }
}