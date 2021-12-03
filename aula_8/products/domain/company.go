package domain

import "context"

type Company struct {
	CNPJ string `json:"cnpj,omitempty"`
}

type CompanyClient interface {
	GetCompany(ctx context.Context, CNPJ string) (*Company, error)
}
