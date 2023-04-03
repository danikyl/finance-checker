package service

import (
	"testing"

	"github.com/danikyl/finance-helper/model"
)

func TestGenerateReport(t *testing.T) {
	financeService := NewFinanceServiceImpl()
	request := model.FinanceReportRequest{
		MonthlyInterestsRate: 0.01,
		RentPrice:            5000,
		InstallmentValue:     8000,
		PeriodInMonths:       120,
	}
	report := financeService.GenerateReport(request)
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
