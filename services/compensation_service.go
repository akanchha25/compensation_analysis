// services/compensation_service.go
package services

import (
	"compensation-api/models"
	"compensation-api/repositories"
	"compensation-api/utils"
	"fmt"
)

type CompensationService interface {
    GetCompensations(filters map[string]interface{}, sortFields map[string]string) ([]models.Compensation, error)
}

type compensationService struct {
    repo repositories.CompensationRepository
}

func NewCompensationService(repo repositories.CompensationRepository) CompensationService {
    return &compensationService{repo: repo}
}

func (s *compensationService) GetCompensations(filters map[string]interface{}, sortFields map[string]string) ([]models.Compensation, error) {
    query := utils.BuildQuery(filters, sortFields)
	fmt.Printf("Elasticsearch query: %s", query.String())

    return s.repo.Search(query)
}
