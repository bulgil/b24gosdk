package crm

import "github.com/bulgil/b24gosdk/transport"

type DealService struct {
	CRMService[Deal]
}

func NewDealService(transport *transport.Transport, webhook string) *DealService {
	return &DealService{
		CRMService: NewCrmService[Deal](transport, webhook, methods{
			add:    "crm.deal.add",
			get:    "crm.deal.get",
			update: "crm.deal.update",
			delete: "crm.deal.delete",
			list:   "crm.deal.list",
		}),
	}
}
