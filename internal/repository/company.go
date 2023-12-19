package repository

import (
	"context"
	"demo-golang/model"
	"fmt"
)

type CompanyRepository interface {
	CreateCompany(ctx context.Context, company model.Company) error
	GetCompany(ctx context.Context, id int) (model.Company, error)
	UpdateCompany(ctx context.Context, company model.Company) error
	DeleteCompany(ctx context.Context, id int) error
	ListCompanies(context.Context, []string) ([]model.Company, error)
}

type companyRepository struct {
}

func NewCompanyRepository() CompanyRepository {
	return &companyRepository{}
}

func (c *companyRepository) CreateCompany(ctx context.Context, company model.Company) error {
	return fmt.Errorf("not implemented")
}

func (c *companyRepository) GetCompany(ctx context.Context, id int) (model.Company, error) {
	return model.Company{
		Id:   1,
		Name: "Test Company",
		Users: []model.User{
			{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@gmail.com",
				Age:       25,
			},
		},
	}, nil
}

func (c *companyRepository) UpdateCompany(ctx context.Context, company model.Company) error {
	return fmt.Errorf("not implemented")
}

func (c *companyRepository) DeleteCompany(ctx context.Context, id int) error {
	return fmt.Errorf("not implemented")
}

func (c *companyRepository) ListCompanies(context.Context, []string) ([]model.Company, error) {
	panic("implement me")
}
