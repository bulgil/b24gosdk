package tasks

import "github.com/bulgil/b24gosdk"

type Task struct {
	ID            *b24gosdk.B24int      `json:"id,omitempty"`
	Title         *string               `json:"title,omitempty"`
	CreatedBy     *b24gosdk.B24int      `json:"createdBy,omitempty"`
	CreatedDate   *b24gosdk.B24datetime `json:"createdDate,omitempty"`
	ResponsibleId *b24gosdk.B24int      `json:"responsibleId,omitempty"`
	UFCRMTask     []*string             `json:"ufCrmTask"`
}
