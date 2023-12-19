package services

import (
	"context"
	"demo-golang/internal/repository"
	"demo-golang/model"
	"fmt"
	"sync"
)

type CompanyService interface {
	CreateCompany(ctx context.Context, company model.Company) error
	GetCompany(ctx context.Context, id int) (model.Company, error)
	UpdateCompany(ctx context.Context, company model.Company) error
	DeleteCompany(ctx context.Context, id int) error
	ListCompanies(ctx context.Context, ids []string) ([]model.Company, error)
}

type companyService struct {
	companyRepository repository.CompanyRepository
}

func NewCompanyService(companyRepository repository.CompanyRepository) CompanyService {
	return &companyService{companyRepository: companyRepository}
}

func (c *companyService) CreateCompany(ctx context.Context, company model.Company) error {
	err := c.companyRepository.CreateCompany(ctx, company)
	if err != nil {
		return fmt.Errorf("error creating company: %w", err)
	}
	return nil
}

func (c *companyService) GetCompany(ctx context.Context, id int) (model.Company, error) {
	company, err := c.companyRepository.GetCompany(ctx, id)
	if err != nil {
		return model.Company{}, fmt.Errorf("error getting company: %w", err)
	}
	return company, nil
}

func (c *companyService) UpdateCompany(ctx context.Context, company model.Company) error {
	err := c.companyRepository.UpdateCompany(ctx, company)
	if err != nil {
		return fmt.Errorf("error updating company: %w", err)
	}
	return nil
}

func (c *companyService) DeleteCompany(ctx context.Context, id int) error {
	err := c.companyRepository.DeleteCompany(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting company: %w", err)
	}
	return nil
}

type Result struct {
	Companies []model.Company
	Err       error
}

func (c *companyService) ListCompanies(ctx context.Context, ids []string) ([]model.Company, error) {
	chunks := c.createChunkOfIds(ids)

	results := make(chan Result, len(chunks))
	var wg sync.WaitGroup

	for _, chunk := range chunks {
		wg.Add(1)
		go func(chunk []string) {
			defer wg.Done()
			companies, err := c.companyRepository.ListCompanies(ctx, chunk)
			results <- Result{Companies: companies, Err: err}
		}(chunk)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var allCompanies []model.Company
	for result := range results {
		if result.Err != nil {
			return nil, fmt.Errorf("error listing companies: %w", result.Err)
		}
		allCompanies = append(allCompanies, result.Companies...)
	}

	return allCompanies, nil
}

func (c *companyService) createChunkOfIds(ids []string) [][]string {
	chunks := make([][]string, 0, (len(ids)+99)/100)

	for i := 0; i < len(ids); i += 100 {
		end := i + 100
		if end > len(ids) {
			end = len(ids)
		}
		chunks = append(chunks, ids[i:end])
	}
	return chunks
}
