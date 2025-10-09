package crm

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Field struct {
	ID         string
	Type       string
	IsReadOnly bool
	IsMultiple bool
	IsDynamic  bool
	Title      string
	Items      []Item
}

type Fields []Field

func (f *Fields) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(data))

	token, err := dec.Token()
	if err != nil {
		return err
	}

	if token != json.Delim('{') {
		return fmt.Errorf("object expected, got: %v", token)
	}

	fields := make([]Field, 0)

	for dec.More() {
		tokenFieldID, err := dec.Token()
		if err != nil {
			return err
		}
		fieldID, ok := tokenFieldID.(string)
		if !ok {
			return fmt.Errorf("expected string key: got %v", tokenFieldID)
		}

		var fieldData struct {
			Type       string `json:"type"`
			IsReadOnly bool   `json:"isReadOnly"`
			IsMultiple bool   `json:"isMultiple"`
			IsDynamic  bool   `json:"isDynamic"`
			Title      string `json:"title"`
			ListLabel  string `json:"listLabel"`
			Items      []Item `json:"items"`
		}

		if err := dec.Decode(&fieldData); err != nil {
			return fmt.Errorf("parsing field %s error: %w", fieldID, err)
		}

		field := Field{
			ID:         fieldID,
			Type:       fieldData.Type,
			IsReadOnly: fieldData.IsReadOnly,
			IsMultiple: fieldData.IsMultiple,
			IsDynamic:  fieldData.IsDynamic,
			Title:      fieldData.Title,
			Items:      fieldData.Items,
		}

		if fieldData.ListLabel != "" {
			field.Title = fieldData.ListLabel
		}

		fields = append(fields, field)
	}

	token, err = dec.Token()
	if err != nil {
		return err
	}
	if token != json.Delim('}') {
		return fmt.Errorf("expected }, got: %v", token)
	}

	*f = fields

	return nil
}

type Item struct {
	ID    string `json:"ID"`
	Value string `json:"VALUE"`
}
