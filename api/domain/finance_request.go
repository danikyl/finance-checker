package domain

type FinanceReportRequest struct {
	MonthlyInterestsRate float64 `json:"monthly_interests_rate"`
	RentPrice            float64 `json:"rent_price"`
	InstallmentValue     float64 `json:"installment_value"`
	PeriodInMonths       int     `json:"period_in_months"`
}
