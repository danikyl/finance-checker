package service

import "github.com/danikyl/finance-helper/model"

type financeServiceImpl struct{}

func (financeService *financeServiceImpl) GenerateReport(request model.FinanceReportRequest) *model.FinanceReportResponse {
	financeResponse := model.FinanceReportResponse{
		SavedMoneyRentInsteadInstallment: getSavedMoneyWhenRentingAndContributingMonthly(request),
		HouseAmountAfterInterests:        getHouseTotalValueAfterInterests(request),
	}

	return &financeResponse
}

func getSavedMoneyWhenRentingAndContributingMonthly(request model.FinanceReportRequest) float64 {
	monthlyContribution := request.InstallmentValue - request.RentPrice
	var finalAmount float64
	for i := 0; i < int(request.PeriodInMonths); i++ {
		finalAmount += monthlyContribution
		finalAmount *= 1.0 + request.MonthlyInterestsRate
	}
	return finalAmount
}

func getHouseTotalValueAfterInterests(request model.FinanceReportRequest) float64 {
	return request.InstallmentValue * float64(request.PeriodInMonths)
}

func NewFinanceServiceImpl() *financeServiceImpl {
	return &financeServiceImpl{}
}
