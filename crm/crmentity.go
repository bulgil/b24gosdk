package crm

import (
	"encoding/json"
	"fmt"

	"github.com/bulgil/b24gosdk"
)

func unmarshalCRMEntity(data []byte, base any, uf *b24gosdk.Userfields) error {
	const op = "UnmarshalCRMEntity"

	if err := json.Unmarshal(data, base); err != nil {
		return fmt.Errorf("%s: base fields unmarshal failed: %w", op, err)
	}

	if err := json.Unmarshal(data, uf); err != nil {
		return fmt.Errorf("%s: userfields unmarshal failed: %w", op, err)
	}

	return nil
}
