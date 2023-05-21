package domain

type FinanceReportResponse struct {
	HouseAmountAfterInterests        float64 `json:"house_amount_after_interests  "`
	SavedMoneyRentInsteadInstallment float64 `json:"saved_money_rent_instead_installment"`
}
