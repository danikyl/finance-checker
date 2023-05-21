package mapper

import (
	apiDomain "github.com/danikyl/finance-helper/api/domain"
	coreDomain "github.com/danikyl/finance-helper/core/domain"
)

func ConvertApiRequestToCoreRequest(request apiDomain.FinanceReportRequest) coreDomain.FinanceReportRequest {
	return coreDomain.FinanceReportRequest{MonthlyInterestsRate: request.MonthlyInterestsRate, RentPrice: request.RentPrice,
	InstallmentValue: request.InstallmentValue, PeriodInMonths: request.PeriodInMonths}
}

func ConvertCoreResponseToApiResponse(coreResponse coreDomain.FinanceReportResponse) apiDomain.FinanceReportResponse {
	return apiDomain.FinanceReportResponse{HouseAmountAfterInterests: coreResponse.HouseAmountAfterInterests, 
	SavedMoneyRentInsteadInstallment: coreResponse.SavedMoneyRentInsteadInstallment}
}