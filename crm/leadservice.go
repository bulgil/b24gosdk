package crm

type LeadService struct {
	CRMService[Lead]
}

func NewLeadService(client client, webhook string) *LeadService {
	return &LeadService{
		CRMService: NewCrmService[Lead](client, webhook, methods{
			add:    "crm.lead.add",
			get:    "crm.lead.get",
			update: "crm.lead.update",
			delete: "crm.lead.delete",
			list:   "crm.lead.list",
		}),
	}
}
