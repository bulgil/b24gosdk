package b24gosdk

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type B24int int

func (i *B24int) UnmarshalJSON(data []byte) error {
	const op = "B24int.UnmarshalJSON"

	var decodedValue any

	err := json.Unmarshal(data, &decodedValue)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	switch v := decodedValue.(type) {
	case string:
		intValue, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("%s: cannot convert string to int: %w", op, err)
		}

		*i = B24int(intValue)
	case float64:
		*i = B24int(int(v))

	default:
		return fmt.Errorf("%s: unsupported JSON type: %T", op, decodedValue)
	}

	return nil
}
