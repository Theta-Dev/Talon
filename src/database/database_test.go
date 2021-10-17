package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRelId(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, getRelId(tt.idParam, tt.sctParam))
			assert.Equal(t, tt.expect > 0, isRelSet(tt.idParam, tt.sctParam))
		})
	}
}
