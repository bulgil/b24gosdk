package crm

import (
	"fmt"

	"github.com/bulgil/b24gosdk"
)

type Deal struct {
	ID                  *b24gosdk.B24int      `json:"ID,omitempty"`
	Title               *string               `json:"TITLE,omitempty"`
	CategoryID          *b24gosdk.B24int      `json:"CATEGORY_ID,omitempty"`
	StageID             *string               `json:"STAGE_ID,omitempty"`
	StageSemanticID     *byte                 `json:"STAGE_SEMANTIC_ID,omitempty"`
	IsNew               *byte                 `json:"IS_NEW,omitempty"`
	IsReccuring         *byte                 `json:"IS_RECURRING,omitempty"`
	IsReturnCustomer    *byte                 `json:"IS_RETURN_CUSTOMER,omitempty"`
	IsRepeatedApproach  *byte                 `json:"IS_REPEATED_APPROACH,omitempty"`
	Probability         *b24gosdk.B24int      `json:"PROBABILITY,omitempty"`
	CurrencyID          *string               `json:"CURRENCY_ID,omitempty"`
	Opportunity         *b24gosdk.B24float    `json:"OPPORTUNITY,omitempty"`
	IsManualOpportunity *byte                 `json:"IS_MANUAL_OPPORTUNITY,omitempty"`
	TaxValue            *b24gosdk.B24float    `json:"TAX_VALUE,omitempty"`
	CompanyID           *b24gosdk.B24int      `json:"COMPANY_ID,omitempty"`
	ContactID           *b24gosdk.B24int      `json:"CONTACT_ID,omitempty"`
	ContactIDs          *[]b24gosdk.B24int    `json:"CONTACT_IDS,omitempty"`
	QuoteID             *b24gosdk.B24int      `json:"QUOTE_ID,omitempty"`
	BeginDate           *b24gosdk.B24date     `json:"BEGIN_DATE,omitempty"`
	CloseDate           *b24gosdk.B24date     `json:"CLOSEDATE,omitempty"`
	Opened              *byte                 `json:"OPENED,omitempty"`
	Closed              *byte                 `json:"CLOSED,omitempty"`
	Comments            *string               `json:"COMMENTS,omitempty"`
	AsignedByID         *b24gosdk.B24int      `json:"ASSIGNED_BY_ID,omitempty"`
	CreatedByID         *b24gosdk.B24int      `json:"CREATED_BY_ID,omitempty"`
	ModifyByID          *b24gosdk.B24int      `json:"MODIFY_BY_ID,omitempty"`
	MovedByID           *b24gosdk.B24int      `json:"MOVED_BY_ID,omitempty"`
	DateCreate          *b24gosdk.B24datetime `json:"DATE_CREATE,omitempty"`
	DateModify          *b24gosdk.B24datetime `json:"DATE_MODIFY,omitempty"`
	MovedTime           *b24gosdk.B24datetime `json:"MOVED_TIME,omitempty"`
	SourceID            *string               `json:"SOURCE_ID,omitempty"`
	SourceDescription   *string               `json:"SOURCE_DESCRIPTION,omitempty"`
	LeadID              *b24gosdk.B24int      `json:"LEAD_ID,omitempty"`
	AdditionalInfo      *string               `json:"ADDITIONAL_INFO,omitempty"`
	LocationID          *string               `json:"LOCATION_ID,omitempty"`
	OriginatorID        *string               `json:"ORIGINATOR_ID,omitempty"`
	OriginID            *string               `json:"ORIGIN_ID,omitempty"`
	UTMSource           *string               `json:"UTM_SOURCE,omitempty"`
	UTMMediom           *string               `json:"UTM_MEDIUM,omitempty"`
	UTMCampaign         *string               `json:"UTM_CAMPAIGN,omitempty"`
	UTMContent          *string               `json:"UTM_CONTENT,omitempty"`
	UTMTerm             *string               `json:"UTM_TERM,omitempty"`
	LastActivityTime    *b24gosdk.B24datetime `json:"LAST_ACTIVITY_TIME,omitempty"`
	LastActivityBy      *b24gosdk.B24int      `json:"LAST_ACTIVITY_BY,omitempty"`

	Userfields b24gosdk.Userfields `json:"-"`
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
