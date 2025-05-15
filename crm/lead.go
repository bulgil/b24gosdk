package crm

import (
	"fmt"

	"github.com/bulgil/b24gosdk"
)

type Lead struct {
	ID                  *b24gosdk.B24int          `json:"ID,omitempty"`
	Title               *string                   `json:"TITLE,omitempty"`
	Honorific           *string                   `json:"HONORIFIC,omitempty"`
	Name                *string                   `json:"NAME,omitempty"`
	SecondName          *string                   `json:"SECOND_NAME,omitempty"`
	LastName            *string                   `json:"LAST_NAME,omitempty"`
	CompanyID           *b24gosdk.B24int          `json:"COMPANY_ID,omitempty"`
	CompanyTitle        *string                   `json:"COMPANY_TITLE,omitempty"`
	ContactID           *b24gosdk.B24int          `json:"CONTACT_ID,omitempty"`
	IsReturnCustomer    *string                   `json:"IS_RETURN_CUSTOMER,omitempty"`
	Birthdate           *b24gosdk.B24datetime     `json:"BIRTHDATE,omitempty"`
	SourceID            *string                   `json:"SOURCE_ID,omitempty"`
	SourceDescription   *string                   `json:"SOURCE_DESCRIPTION,omitempty"`
	StatusID            *string                   `json:"STATUS_ID,omitempty"`
	StatusDescription   *string                   `json:"STATUS_DESCRIPTION,omitempty"`
	Post                *string                   `json:"POST,omitempty"`
	Comments            *string                   `json:"COMMENTS,omitempty"`
	CurrencyID          *string                   `json:"CURRENCY_ID,omitempty"`
	Opportunity         *b24gosdk.B24float        `json:"OPPORTUNITY,omitempty"`
	IsManualOpportunity *string                   `json:"IS_MANUAL_OPPORTUNITY,omitempty"`
	HasPhone            *string                   `json:"HAS_PHONE,omitempty"`
	HasEmail            *string                   `json:"HAS_EMAIL,omitempty"`
	HasImol             *string                   `json:"HHAS_IMOL,omitempty"`
	AsignedByID         *b24gosdk.B24int          `json:"ASSIGNED_BY_ID,omitempty"`
	CreatedByID         *b24gosdk.B24int          `json:"CREATED_BY_ID,omitempty"`
	ModifyByID          *b24gosdk.B24int          `json:"MODIFY_BY_ID,omitempty"`
	MovedById           *b24gosdk.B24int          `json:"MOVED_BY_ID,omitempty"`
	DateCreate          *b24gosdk.B24datetime     `json:"DATE_CREATE,omitempty"`
	DateModify          *b24gosdk.B24datetime     `json:"DATE_MODIFY,omitempty"`
	DateClosed          *b24gosdk.B24datetime     `json:"DATE_CLOSED,omitempty"`
	StatusSemanticID    *string                   `json:"STATUS_SEMANTIC_ID,omitempty"`
	Opened              *string                   `json:"OPENED,omitempty"`
	OriginatorID        *string                   `json:"ORIGINATOR_ID,omitempty"`
	OriginID            *string                   `json:"ORIGIN_ID,omitempty"`
	MovedTime           *b24gosdk.B24datetime     `json:"MOVED_TIME,omitempty"`
	Address             *string                   `json:"ADDRESS,omitempty"`
	Address2            *string                   `json:"ADDRESS_2,omitempty"`
	AddressCity         *string                   `json:"ADDRESS_CITY,omitempty"`
	AddressPostalCode   *string                   `json:"ADDRESS_POSTAL_CODE,omitempty"`
	AddressRegion       *string                   `json:"ADDRESS_REGION,omitempty"`
	AddressProvince     *string                   `json:"ADDRESS_PROVINCE,omitempty"`
	AddressCountry      *string                   `json:"ADDRESS_COUNTRY,omitempty"`
	AddressCountryCode  *string                   `json:"ADDRESS_COUNTRY_CODE,omitempty"`
	AddressLocAddrID    *string                   `json:"ADDRESS_LOC_ADDR_ID,omitempty"`
	URMSource           *string                   `json:"UTM_SOURCE,omitempty"`
	UTMMedium           *string                   `json:"UTM_MEDIUM,omitempty"`
	UTMCampaign         *string                   `json:"URM_CAMPAIGN,omitempty"`
	UTMContent          *string                   `json:"UTM_CONTENT,omitempty"`
	UTMTerm             *string                   `json:"UTM_TERM,omitempty"`
	LastActivityBy      *b24gosdk.B24int          `json:"LAST_ACTIVITY_BY,omitempty"`
	LastActivityTime    *b24gosdk.B24datetime     `json:"LAST_ACTIVITY_TIME,omitempty"`
	Phone               *[]b24gosdk.B24multifield `json:"PHONE,omitempty"`
	IM                  *[]b24gosdk.B24multifield `json:"IM,omitempty"`

	Userfields b24gosdk.Userfields `json:"-"`
}

func (d *Lead) UnmarshalJSON(data []byte) error {
	const op = "Lead.UnmarshalJSON"

	type Alias Lead
	err := unmarshalCRMEntity(data, (*Alias)(d), &d.Userfields)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
