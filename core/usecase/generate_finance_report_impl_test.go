package usecase

import (
	"testing"

	"github.com/danikyl/finance-helper/core/domain"
)

func TestGenerateReport(t *testing.T) {
	generateFinanceReportUseCase := NewGenerateFinanceReportImpl()
	request := domain.FinanceReportRequest{
		MonthlyInterestsRate: 0.01,
		RentPrice:            5000,
		InstallmentValue:     8000,
		PeriodInMonths:       120,
	}
	report := generateFinanceReportUseCase.Execute(request)
	got := report.HouseAmountAfterInterests
	want := float64(960000)
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
	got = report.SavedMoneyRentInsteadInstallment
	want = 697017.2290558208
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
