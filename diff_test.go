package ormdiff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Name struct {
	first string
	last  string
	len   int
}

func TestPatch(t *testing.T) {
	name1 := Name{
		first: "abc",
		len:   24,
	}
	name2 := Name{
		first: "ccc",
		last:  "bbb",
		len:   34,
	}
	differ := OrmDiff{}
	differ.Up(name1)
	assert.Equal(t, differ.GetUpFields(), map[string]interface{}{
		"first": "abc",
		"len":   24,
	})

	differ.Up(name2)
	assert.Equal(t, differ.GetUpFields(), map[string]interface{}{
		"first": "ccc",
		"last":  "bbb",
		"len":   34,
	})
}
