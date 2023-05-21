package domain

type FinanceReportRequest struct {
	MonthlyInterestsRate float64
	RentPrice            float64
	InstallmentValue     float64
	PeriodInMonths       int
}
