package service

import "github.com/danikyl/finance-helper/model"

type financeServiceImpl struct{}

func (financeService *financeServiceImpl) GenerateReport(request model.FinanceReportRequest) *model.FinanceReportResponse {
	financeResponse := model.FinanceReportResponse{}

	return &financeResponse
}

func NewFinanceServiceImpl() *financeServiceImpl {
	return &financeServiceImpl{}
}
