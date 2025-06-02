package crm

type LeadService struct {
	CRMService[Lead]
}

func NewLeadService(transport Transport, webhook string) *LeadService {
	return &LeadService{
		CRMService: NewCrmService[Lead](transport, webhook, methods{
			add:    "crm.lead.add",
			get:    "crm.lead.get",
			update: "crm.lead.update",
			delete: "crm.lead.delete",
			list:   "crm.lead.list",
		}),
	}
}
