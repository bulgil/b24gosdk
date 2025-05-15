package b24gosdk

import (
	"encoding/json"
	"testing"
)

func Test_B24intUnmarshalJSON(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected B24int
		wantErr  bool
	}{
		{
			name:     "valid string number",
			input:    `"42"`,
			expected: 42,
			wantErr:  false,
		},

		{
			name:     "valid numeric number",
			input:    `42`,
			expected: 42,
			wantErr:  false,
		},

		{
			name:    "invalid string",
			input:   "abc",
			wantErr: true,
		},

		{
			name:    "null input",
			input:   `null`,
			wantErr: true,
		},

		{
			name:    "unsupported type (boolean)",
			input:   `true`,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var b B24int
			err := json.Unmarshal([]byte(c.input), &b)

			if (err != nil) != c.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, c.wantErr)
			}

			if !c.wantErr && b != c.expected {
				t.Errorf("UnmarshalJSON() got = %v, want %v", b, c.expected)
			}
		})
	}
}
