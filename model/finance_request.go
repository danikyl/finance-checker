package model

type FinanceReportRequest struct {
	CurrentCdi       float32 `json:"current_cdi"`
	TotalInstallment float32 `json:"total_installment"`
	InstallmentInterests   float32 `json:"installment_interests"`
}
