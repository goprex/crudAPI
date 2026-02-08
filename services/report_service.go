package services

import (
	"crudapi/models"
	"crudapi/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetTodayReport() (*models.DailyReport, error) {
	return s.repo.GetTodayReport()
}

