package handlers

import "car-rent/internal/business"

type Handlers struct {
}

func NewHandler(business business.Business) Handlers {
	return Handlers{}
}
