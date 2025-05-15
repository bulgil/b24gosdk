package crm

import (
	"fmt"

	"github.com/bulgil/b24gosdk"
)

type Contact struct {
	ID         *b24gosdk.B24int `json:"ID"`
	Post       *string          `json:"POST"`
	Comments   *string          `json:"COMMENTS"`
	Honorific  *string          `json:"HONORIFIC"`
	Name       *string          `json:"NAME"`
	SecondName *string          `json:"SECOND_NAME"`
	LastName   *string          `json:"LAST_NAME"`
	// Photo
	LeadID            *b24gosdk.B24int      `json:"LEAD_ID"`
	TypeID            *string               `json:"TYPE_ID"`
	SourceID          *string               `json:"SOURCE_ID"`
	SourceDescription *string               `json:"SOURCE_DESCRIPTION"`
	CompanyID         *string               `json:"COMPANY_ID"`
	Birthday          *b24gosdk.B24datetime `json:"BIRTHDAY"`
	Export            *string               `json:"EXPORT"`
	HasPhone          *string               `json:"HAS_PHONE"`
	HasEmail          *string               `json:"HAS_EMAIL"`
	HasImol           *string               `json:"HAS_IMOL"`
	DateCreate        *b24gosdk.B24datetime `json:"DATE_CREATE"`
	DateModify        *b24gosdk.B24datetime `json:"DATE_MODIFY"`
	AssignedByID      *b24gosdk.B24int      `json:"ASSIGNED_BY_ID"`
	CreateByID        *b24gosdk.B24int      `json:"CREATED_BY_ID"`
	ModifyByID        *b24gosdk.B24int      `json:"MODIFY_BY_ID"`
	Opened            *string               `json:"OPENED"`
	// FaceID
	LastActivityTime *b24gosdk.B24datetime     `json:"LAST_ACTIVITY_TIME"`
	LastActivityBy   *b24gosdk.B24int          `json:"LAST_ACTIVITY_BY"`
	UTMSource        *string                   `json:"UTM_SOURCE"`
	UTMMedium        *string                   `json:"UTM_MEDIUM"`
	UTMCampaign      *string                   `json:"UTM_CAMPAIGN"`
	UTMContent       *string                   `json:"UTM_CONTENT"`
	UTMTerm          *string                   `json:"UTM_TERM"`
	Phone            *[]b24gosdk.B24multifield `json:"PHONE"`
	Email            *[]b24gosdk.B24multifield `json:"EMAIL"`
	Web              *[]b24gosdk.B24multifield `json:"WEB"`
	IM               *[]b24gosdk.B24multifield `json:"IM"`
	Link             *[]b24gosdk.B24multifield `json:"LINK"`
	OriginatorID     *string                   `json:"ORIGINATOR_ID"`
	OriginID         *string                   `json:"ORIGIN_ID"`
	OriginVersion    *string                   `json:"ORIGIN_VERSION"`

	Userfields b24gosdk.Userfields `json:"-"`
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
