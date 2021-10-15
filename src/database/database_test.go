package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRelId(t *testing.T) {
	params := []struct {
		name     string
		idParam  uint
		sctParam interface{}
		expect   uint
	}{
		{
			name:    "id_set",
			idParam: 10,
			expect:  10,
		},
		{
			name:     "sct_set",
			sctParam: Website{ID: 1},
			expect:   1,
		},
		{
			name:     "sct_ptr_set",
			sctParam: &Website{ID: 5},
			expect:   5,
		},
		{
			name:   "empty",
			expect: 0,
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			assert.Equal(t, p.expect, getRelId(p.idParam, p.sctParam))
			assert.Equal(t, p.expect > 0, isRelSet(p.idParam, p.sctParam))
		})
	}
}
