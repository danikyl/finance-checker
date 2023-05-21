package usecase

import "github.com/danikyl/finance-helper/core/domain"

type generateFinanceReportImpl struct{}

func (*generateFinanceReportImpl) Execute(request domain.FinanceReportRequest) *domain.FinanceReportResponse {
	financeResponse := domain.FinanceReportResponse{
		SavedMoneyRentInsteadInstallment: getSavedMoneyWhenRentingAndContributingMonthly(request),
		HouseAmountAfterInterests:        getHouseTotalValueAfterInterests(request),
	}

	return &financeResponse
}

func getSavedMoneyWhenRentingAndContributingMonthly(request domain.FinanceReportRequest) float64 {
	monthlyContribution := request.InstallmentValue - request.RentPrice
	var finalAmount float64
	for i := 0; i < int(request.PeriodInMonths); i++ {
		finalAmount += monthlyContribution
		finalAmount *= 1.0 + request.MonthlyInterestsRate
	}
	return finalAmount
}

func getHouseTotalValueAfterInterests(request domain.FinanceReportRequest) float64 {
	return request.InstallmentValue * float64(request.PeriodInMonths)
}

func NewGenerateFinanceReportImpl() *generateFinanceReportImpl {
	return &generateFinanceReportImpl{}
}
