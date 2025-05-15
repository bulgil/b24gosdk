package crm

import (
	"fmt"

	"github.com/bulgil/b24gosdk"
)

type Deal struct {
	ID          b24gosdk.B24int     `json:"ID"`
	Title       string              `json:"TITLE"`
	AsignedByID b24gosdk.B24int     `json:"ASSIGNED_BY_ID"`
	Userfields  b24gosdk.Userfields `json:"-"`
}

func (d *Deal) UnmarshalJSON(data []byte) error {
	const op = "Deal.UnmarshalJSON"

	type Alias Deal
	err := unmarshalCRMEntity(data, (*Alias)(d), &d.Userfields)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
