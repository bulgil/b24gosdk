package b24gosdk

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
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

type B24float float64

func (i *B24float) UnmarshalJSON(data []byte) error {
	const op = "B24float.UnmarshalJSON"

	var decodedValue any

	err := json.Unmarshal(data, &decodedValue)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	switch v := decodedValue.(type) {
	case string:
		intValue, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return fmt.Errorf("%s: cannot convert string to float64: %w", op, err)
		}

		*i = B24float(intValue)
	case float64:
		*i = B24float(v)

	default:
		return fmt.Errorf("%s: unsupported JSON type: %T", op, decodedValue)
	}

	return nil
}

type B24date time.Time

func (d *B24date) UnmarshalJSON(data []byte) error {
	const op = "B24date.UnmarshalJSON"
	const dateLayout = "2006-01-02"

	var raw string
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	t, err := time.Parse(dateLayout, raw)
	if err != nil {
		return fmt.Errorf("%s: time parse error :%w", op, err)
	}

	*d = B24date(t)

	return nil
}

type B24datetime time.Time

func (d *B24datetime) UnmarshalJSON(data []byte) error {
	const op = "B24datetime.UnmarshalJSON"
	const datetimeLayout = "2006-01-02T15:04:05"

	var raw string
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	t, err := time.Parse(datetimeLayout, raw)
	if err != nil {
		return fmt.Errorf("%s: time parse error :%w", op, err)
	}

	*d = B24datetime(t)

	return nil
}
