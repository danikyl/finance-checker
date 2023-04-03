package service

import "github.com/danikyl/finance-helper/model"

type FinanceService interface {
	GenerateReport(model.FinanceReportRequest) *model.FinanceReportResponse
}
