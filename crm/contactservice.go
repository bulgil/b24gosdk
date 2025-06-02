package crm

type ContactService struct {
	CRMService[Contact]
}

func NewContactService(transport Transport, webhook string) *ContactService {
	return &ContactService{
		CRMService: NewCrmService[Contact](transport, webhook, methods{
			add:    "crm.contact.add",
			get:    "crm.contact.get",
			update: "crm.contact.update",
			delete: "crm.contact.delete",
			list:   "crm.contact.list",
		}),
	}
}
