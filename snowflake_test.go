package snowflake

import (
	"encoding/json"
	"errors"
	"strconv"
	"testing"
)

func TestID_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name          string
		data          []byte
		expected      ID
		allowUnquoted bool
		err           error
	}{
		{
			name:          "null",
			data:          []byte("null"),
			expected:      0,
			allowUnquoted: false,
			err:           nil,
		},
		{
			name:          "valid id",
			data:          []byte(`"123456"`),
			expected:      123456,
			allowUnquoted: false,
			err:           nil,
		},
		{
			name:          "invalid id",
			data:          []byte(`"id"`),
			expected:      0,
			allowUnquoted: false,
			err:           strconv.ErrSyntax,
		},
		{
			name:          "unquoted 0",
			data:          []byte("0"),
			expected:      0,
			allowUnquoted: false,
			err:           nil,
		},
		{
			name:          "quoted 0",
			data:          []byte(`"0"`),
			expected:      0,
			allowUnquoted: false,
			err:           nil,
		},
		{
			name:          "unquoted 1 no allowUnquoted",
			data:          []byte("1"),
			expected:      0,
			allowUnquoted: false,
			err:           strconv.ErrSyntax,
		},
		{
			name:          "unquoted 1 allowUnquoted",
			data:          []byte("1"),
			expected:      1,
			allowUnquoted: true,
			err:           nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AllowUnquoted = tt.allowUnquoted
			var id ID
			err := json.Unmarshal(tt.data, &id)

			if !errors.Is(err, tt.err) {
				t.Errorf("expected error %v, got %v", tt.err, err)
			}
			if id != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, id)
			}
		})
	}
}
