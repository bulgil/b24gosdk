package crm

import (
	"fmt"

	"github.com/bulgil/b24gosdk"
)

type Contact struct {
	ID                *b24gosdk.B24int          `json:"ID,omitempty"`
	Post              *string                   `json:"POST,omitempty"`
	Comments          *string                   `json:"COMMENTS,omitempty"`
	Honorific         *string                   `json:"HONORIFIC,omitempty"`
	Name              *string                   `json:"NAME,omitempty"`
	SecondName        *string                   `json:"SECOND_NAME,omitempty"`
	LastName          *string                   `json:"LAST_NAME,omitempty"`
	LeadID            *b24gosdk.B24int          `json:"LEAD_ID,omitempty"`
	TypeID            *string                   `json:"TYPE_ID,omitempty"`
	SourceID          *string                   `json:"SOURCE_ID,omitempty"`
	SourceDescription *string                   `json:"SOURCE_DESCRIPTION,omitempty"`
	CompanyID         *string                   `json:"COMPANY_ID,omitempty"`
	Birthday          *b24gosdk.B24datetime     `json:"BIRTHDAY,omitempty"`
	Export            *string                   `json:"EXPORT,omitempty"`
	HasPhone          *string                   `json:"HAS_PHONE,omitempty"`
	HasEmail          *string                   `json:"HAS_EMAIL,omitempty"`
	HasImol           *string                   `json:"HAS_IMOL,omitempty"`
	DateCreate        *b24gosdk.B24datetime     `json:"DATE_CREATE,omitempty"`
	DateModify        *b24gosdk.B24datetime     `json:"DATE_MODIFY,omitempty"`
	AssignedByID      *b24gosdk.B24int          `json:"ASSIGNED_BY_ID,omitempty"`
	CreateByID        *b24gosdk.B24int          `json:"CREATED_BY_ID,omitempty"`
	ModifyByID        *b24gosdk.B24int          `json:"MODIFY_BY_ID,omitempty"`
	Opened            *string                   `json:"OPENED,omitempty"`
	LastActivityTime  *b24gosdk.B24datetime     `json:"LAST_ACTIVITY_TIME,omitempty"`
	LastActivityBy    *b24gosdk.B24int          `json:"LAST_ACTIVITY_BY,omitempty"`
	UTMSource         *string                   `json:"UTM_SOURCE,omitempty"`
	UTMMedium         *string                   `json:"UTM_MEDIUM,omitempty"`
	UTMCampaign       *string                   `json:"UTM_CAMPAIGN,omitempty"`
	UTMContent        *string                   `json:"UTM_CONTENT,omitempty"`
	UTMTerm           *string                   `json:"UTM_TERM,omitempty"`
	Phone             *[]b24gosdk.B24multifield `json:"PHONE,omitempty"`
	Email             *[]b24gosdk.B24multifield `json:"EMAIL,omitempty"`
	Web               *[]b24gosdk.B24multifield `json:"WEB,omitempty"`
	IM                *[]b24gosdk.B24multifield `json:"IM,omitempty"`
	Link              *[]b24gosdk.B24multifield `json:"LINK,omitempty"`
	OriginatorID      *string                   `json:"ORIGINATOR_ID,omitempty"`
	OriginID          *string                   `json:"ORIGIN_ID,omitempty"`
	OriginVersion     *string                   `json:"ORIGIN_VERSION,omitempty"`

	// Photo
	// FaceID

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
