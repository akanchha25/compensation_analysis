// services/compensation_service.go
package services

import (
    "compensation-api/repositories"
    "compensation-api/models"
    "compensation-api/utils"
)

type CompensationService interface {
    GetCompensations(filters map[string]string, sortFields map[string]string) ([]models.Compensation, error)
}

type compensationService struct {
    repo repositories.CompensationRepository
}

func NewCompensationService(repo repositories.CompensationRepository) CompensationService {
    return &compensationService{repo: repo}
}

func (s *compensationService) GetCompensations(filters map[string]string, sortFields map[string]string) ([]models.Compensation, error) {
    query := utils.BuildQuery(filters, sortFields)
    return s.repo.Search(query)
}
