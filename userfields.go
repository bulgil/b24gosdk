package b24gosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Userfields map[string]any

func (uf *Userfields) UnmarshalJSON(data []byte) error {
	const op = "Userfields.UnmarshalJSON"

	*uf = make(Userfields)

	dec := json.NewDecoder(bytes.NewReader(data))

	t, err := dec.Token()
	if err != nil {
		return fmt.Errorf("%s: error with token get: %w", op, err)
	}

	if delim, ok := t.(json.Delim); !ok || delim != '{' {
		return fmt.Errorf("%s: expected JSON object", op)
	}

	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		key, ok := t.(string)
		if !ok {
			return fmt.Errorf("%s: expected string token, got %T", op, key)
		}

		if strings.HasPrefix(key, "UF_CRM_") {
			var val any
			if err := dec.Decode(&val); err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
			(*uf)[key] = val
		} else {
			var skip json.RawMessage
			if err := dec.Decode(&skip); err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
		}
	}

	_, err = dec.Token()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
