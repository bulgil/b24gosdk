package b24gosdk

import (
	"encoding/json"
	"testing"
	"time"
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

func TestB24date_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantDate  time.Time
		wantError bool
	}{
		{
			name:     "valid date",
			input:    `"2024-05-15"`,
			wantDate: time.Date(2024, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "invalid format",
			input:     `"15-05-2024"`,
			wantError: true,
		},
		{
			name:      "null value",
			input:     `null`,
			wantError: true,
		},
		{
			name:      "non-string JSON value",
			input:     `12345`,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d B24date
			err := json.Unmarshal([]byte(tt.input), &d)

			if (err != nil) != tt.wantError {
				t.Fatalf("expected error = %v, got = %v", tt.wantError, err)
			}

			if !tt.wantError {
				got := time.Time(d)
				if !got.Equal(tt.wantDate) {
					t.Errorf("expected date = %v, got = %v", tt.wantDate, got)
				}
			}
		})
	}
}

func TestB24datetime_UnmarshalJSON_Valid(t *testing.T) {
	const layout = "2006-01-02T15:04:05"
	input := `"2025-05-15T14:30:00"`
	var dt B24datetime

	err := json.Unmarshal([]byte(input), &dt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parsedTime, _ := time.Parse(layout, "2025-05-15T14:30:00")
	if time.Time(dt) != parsedTime {
		t.Errorf("expected %v, got %v", parsedTime, dt)
	}
}

func TestB24datetime_UnmarshalJSON_EmptyString(t *testing.T) {
	input := `""`
	var dt B24datetime

	err := json.Unmarshal([]byte(input), &dt)
	if err == nil {
		t.Fatalf("expected error for empty string, got nil")
	}
}

func TestB24datetime_UnmarshalJSON_InvalidFormat(t *testing.T) {
	input := `"15-05-2025 14:30:00"`
	var dt B24datetime

	err := json.Unmarshal([]byte(input), &dt)
	if err == nil {
		t.Fatalf("expected error for invalid format, got nil")
	}
}
