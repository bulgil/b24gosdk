package b24gosdk

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_UserfieldsUnmarshalJSON(t *testing.T) {
	var cases = []struct {
		name          string
		input         string
		expected      Userfields
		isErrExpected bool
	}{
		{name: "valid json",
			input: `{
			"ID": 256,
			"TITLE": "testing title",
			"UF_CRM_1": "somethign",
			"UF_CRM_2": "something 2"
		}`,
			expected: Userfields{
				"UF_CRM_1": "somethign",
				"UF_CRM_2": "something 2",
			},
			isErrExpected: false,
		},

		{
			name: "empty userfields",
			input: `{
			"ID": 245,
			"TITLE": "testing title"
		}`,
			expected:      Userfields{},
			isErrExpected: false,
		},

		{
			name: "only userfields",
			input: `{
			"UF_CRM_1": 11233321,
			"UF_CRM_2": "something 2"
		}`,
			expected: Userfields{
				"UF_CRM_1": float64(11233321),
				"UF_CRM_2": "something 2",
			},
			isErrExpected: false,
		},

		{
			name:          "invalid json",
			input:         `{`,
			expected:      Userfields{},
			isErrExpected: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var uf Userfields

			err := json.Unmarshal([]byte(c.input), &uf)

			if c.isErrExpected && err == nil {
				t.Fatalf("error expected, got no error")
			}

			if c.isErrExpected && err != nil {
				t.Logf("expected error received: %v", err)
				return
			}

			if !c.isErrExpected && err != nil {
				t.Fatalf("error is not expected, got %s", err)
			}

			if !reflect.DeepEqual(uf, c.expected) {
				t.Fatalf("expected %v, got %v", c.expected, uf)
			}
		})
	}
}
