package crm

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/bulgil/b24gosdk"
)

func Test_DealUnmarshalJSON(t *testing.T) {
	cases := []struct {
		name      string
		input     string
		expected  Deal
		expectErr bool
	}{
		{
			name: "valid deal with userfields",
			input: `{
				"ID": 12345,
				"TITLE": "Some title",
				"ASSIGNED_BY_ID": 156,
				"UF_CRM_1": 123,
				"UF_CRM_22": "somesome"
			}`,
			expected: Deal{
				ID:          12345,
				Title:       "Some title",
				AsignedByID: 156,
				Userfields: b24gosdk.Userfields{
					"UF_CRM_1":  float64(123),
					"UF_CRM_22": "somesome",
				},
			},
		},
		{
			name: "deal without userfields",
			input: `{
				"ID": 123,
				"TITLE": "No userfields",
				"ASSIGNED_BY_ID": 456
			}`,
			expected: Deal{
				ID:          123,
				Title:       "No userfields",
				AsignedByID: 456,
				Userfields:  b24gosdk.Userfields{},
			},
		},
		{
			name: "userfield with null value",
			input: `{
				"ID": 555,
				"TITLE": "Null field test",
				"ASSIGNED_BY_ID": 888,
				"UF_CRM_NULL": null
			}`,
			expected: Deal{
				ID:          555,
				Title:       "Null field test",
				AsignedByID: 888,
				Userfields: b24gosdk.Userfields{
					"UF_CRM_NULL": nil,
				},
			},
		},
		{
			name: "userfield with array value",
			input: `{
				"ID": 777,
				"TITLE": "Array field test",
				"ASSIGNED_BY_ID": 999,
				"UF_CRM_LIST": [1, 2, 3]
			}`,
			expected: Deal{
				ID:          777,
				Title:       "Array field test",
				AsignedByID: 999,
				Userfields: b24gosdk.Userfields{
					"UF_CRM_LIST": []interface{}{float64(1), float64(2), float64(3)},
				},
			},
		},
		{
			name:      "invalid json",
			input:     `{ "ID": 1, "TITLE": `,
			expectErr: true,
		},
		{
			name: "userfield with object",
			input: `{
				"ID": 987,
				"TITLE": "Object field",
				"ASSIGNED_BY_ID": 654,
				"UF_CRM_OBJ": { "foo": "bar" }
			}`,
			expected: Deal{
				ID:          987,
				Title:       "Object field",
				AsignedByID: 654,
				Userfields: b24gosdk.Userfields{
					"UF_CRM_OBJ": map[string]interface{}{
						"foo": "bar",
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var deal Deal
			err := json.Unmarshal([]byte(c.input), &deal)

			if c.expectErr {
				if err == nil {
					t.Fatalf("expected error, got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(deal, c.expected) {
				t.Fatalf("expected %+v, got %+v", c.expected, deal)
			}
		})
	}
}
