package crm

type DealService struct {
	CRMService[Deal]
}

func NewDealService(webhook string) *DealService {
	return &DealService{
		CRMService: NewCrmService[Deal](webhook, methods{
			add:    "crm.deal.add",
			get:    "crm.deal.get",
			update: "crm.deal.update",
			delete: "crm.deal.delete",
			list:   "crm.deal.list",
		}),
	}
}
