package usecase

import "github.com/danikyl/finance-helper/core/domain"

type GenerateFinanceReport interface {
	Execute(domain.FinanceReportRequest) *domain.FinanceReportResponse
}
