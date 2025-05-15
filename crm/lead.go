package crm

import (
	"fmt"

	"github.com/bulgil/b24gosdk"
)

type Lead struct {
	ID                  *b24gosdk.B24int          `json:"ID"`
	Title               *string                   `json:"TITLE"`
	Honorific           *string                   `json:"HONORIFIC"`
	Name                *string                   `json:"NAME"`
	SecondName          *string                   `json:"SECOND_NAME"`
	LastName            *string                   `json:"LAST_NAME"`
	CompanyID           *b24gosdk.B24int          `json:"COMPANY_ID"`
	CompanyTitle        *string                   `json:"COMPANY_TITLE"`
	ContactID           *b24gosdk.B24int          `json:"CONTACT_ID"`
	IsReturnCustomer    *string                   `json:"IS_RETURN_CUSTOMER"`
	Birthdate           *b24gosdk.B24datetime     `json:"BIRTHDATE"`
	SourceID            *string                   `json:"SOURCE_ID"`
	SourceDescription   *string                   `json:"SOURCE_DESCRIPTION"`
	StatusID            *string                   `json:"STATUS_ID"`
	StatusDescription   *string                   `json:"STATUS_DESCRIPTION"`
	Post                *string                   `json:"POST"`
	Comments            *string                   `json:"COMMENTS"`
	CurrencyID          *string                   `json:"CURRENCY_ID"`
	Opportunity         *b24gosdk.B24float        `json:"OPPORTUNITY"`
	IsManualOpportunity *string                   `json:"IS_MANUAL_OPPORTUNITY"`
	HasPhone            *string                   `json:"HAS_PHONE"`
	HasEmail            *string                   `json:"HAS_EMAIL"`
	HasImol             *string                   `json:"HHAS_IMOL"`
	AsignedByID         *b24gosdk.B24int          `json:"ASSIGNED_BY_ID"`
	CreatedByID         *b24gosdk.B24int          `json:"CREATED_BY_ID"`
	ModifyByID          *b24gosdk.B24int          `json:"MODIFY_BY_ID"`
	MovedById           *b24gosdk.B24int          `json:"MOVED_BY_ID"`
	DateCreate          *b24gosdk.B24datetime     `json:"DATE_CREATE"`
	DateModify          *b24gosdk.B24datetime     `json:"DATE_MODIFY"`
	DateClosed          *b24gosdk.B24datetime     `json:"DATE_CLOSED"`
	StatusSemanticID    *string                   `json:"STATUS_SEMANTIC_ID"`
	Opened              *string                   `json:"OPENED"`
	OriginatorID        *string                   `json:"ORIGINATOR_ID"`
	OriginID            *string                   `json:"ORIGIN_ID"`
	MovedTime           *b24gosdk.B24datetime     `json:"MOVED_TIME"`
	Address             *string                   `json:"ADDRESS"`
	Address2            *string                   `json:"ADDRESS_2"`
	AddressCity         *string                   `json:"ADDRESS_CITY"`
	AddressPostalCode   *string                   `json:"ADDRESS_POSTAL_CODE"`
	AddressRegion       *string                   `json:"ADDRESS_REGION"`
	AddressProvince     *string                   `json:"ADDRESS_PROVINCE"`
	AddressCountry      *string                   `json:"ADDRESS_COUNTRY"`
	AddressCountryCode  *string                   `json:"ADDRESS_COUNTRY_CODE"`
	AddressLocAddrID    *string                   `json:"ADDRESS_LOC_ADDR_ID"`
	URMSource           *string                   `json:"UTM_SOURCE"`
	UTMMedium           *string                   `json:"UTM_MEDIUM"`
	UTMCampaign         *string                   `json:"URM_CAMPAIGN"`
	UTMContent          *string                   `json:"UTM_CONTENT"`
	UTMTerm             *string                   `json:"UTM_TERM"`
	LastActivityBy      *b24gosdk.B24int          `json:"LAST_ACTIVITY_BY"`
	LastActivityTime    *b24gosdk.B24datetime     `json:"LAST_ACTIVITY_TIME"`
	Phone               *[]b24gosdk.B24multifield `json:"PHONE"`
	IM                  *[]b24gosdk.B24multifield `json:"IM"`

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
