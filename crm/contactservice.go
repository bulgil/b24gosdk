package crm

type ContactService struct {
	CRMService[Contact]
}

func NewContactService(client client, webhook string) *ContactService {
	return &ContactService{
		CRMService: NewCrmService[Contact](client, webhook, methods{
			add:    "crm.contact.add",
			get:    "crm.contact.get",
			update: "crm.contact.update",
			delete: "crm.contact.delete",
			list:   "crm.contact.list",
		}),
	}
}
