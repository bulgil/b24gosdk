package crm

import "github.com/bulgil/b24gosdk"

type DealService struct {
	CRMService[Deal]
}

func NewDealService(webhook string) *DealService {
	return &DealService{
		CRMService: CRMService[Deal]{
			client:  b24gosdk.NewClient(nil, webhook),
			webhook: webhook,
			methods: methods{
				add:    "crm.deal.add",
				get:    "crm.deal.get",
				update: "crm.deal.update",
				delete: "crm.deal.delete",
				list:   "crm.deal.list",
			},
		},
	}
}
