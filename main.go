package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type financeReportRequest struct {
	CurrentCdi       float32 `json:"current_cdi"`
	TotalInstallment float32 `json:"total_installment"`
	InstallmentFee   float32 `json:"installment_fee"`
}

type financeReportResponse struct {
	HouseAmountBeforeFees            float32 `json:"house_amount_before_fees"`
	HouseAmountAfterFees             float32 `json:"house_amount_after_fees"`
	RentSpentAmount                  float32 `json:"rent_spent_amount"`
	SavedMoneyRentInsteadInstallment float32 `json:"saved_money_rent_instead_installment"`
	InvestedSavedMoneyTotalAmount    float32 `json:"invested_saved_money_total_amount"`
}

func generateReport(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, financeReportResponse{})
}

func main() {
	router := gin.Default()
	router.GET("/albums", generateReport)

	router.Run("localhost:8080")
}
