package provider

import (
	"car-rent/bootstrap"
)

type Provider struct {
}

func NewProvider(cfg bootstrap.Providers) Provider {
	return Provider{}
}
