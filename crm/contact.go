package crm

import (
	"fmt"

	"github.com/bulgil/b24gosdk"
)

type Contact struct {
	ID          b24gosdk.B24int     `json:"ID"`
	Title       string              `json:"TITLE"`
	AsignedByID b24gosdk.B24int     `json:"ASSIGNED_BY_ID"`
	Userfields  b24gosdk.Userfields `json:"-"`
}

func (d *Contact) UnmarshalJSON(data []byte) error {
	const op = "Contact.UnmarshalJSON"

	type Alias Contact
	err := unmarshalCRMEntity(data, (*Alias)(d), &d.Userfields)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
