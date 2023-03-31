package model

type FinanceReportResponse struct {
	HouseAmountBeforeFees            float32 `json:"house_amount_before_fees"`
	HouseAmountAfterFees             float32 `json:"house_amount_after_fees"`
	RentSpentAmount                  float32 `json:"rent_spent_amount"`
	SavedMoneyRentInsteadInstallment float32 `json:"saved_money_rent_instead_installment"`
	InvestedSavedMoneyTotalAmount    float32 `json:"invested_saved_money_total_amount"`
}
