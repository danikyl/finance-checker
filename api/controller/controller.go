package controller

import (
	"net/http"

	"github.com/danikyl/finance-helper/api/domain"
	"github.com/danikyl/finance-helper/api/mapper"
	"github.com/danikyl/finance-helper/core/usecase"
	"github.com/gin-gonic/gin"
)

type ApiController struct {
	generateFinanceReportUsecase usecase.GenerateFinanceReport
}

func NewController(generateFinanceReportUsecase usecase.GenerateFinanceReport) *ApiController {
	return &ApiController{generateFinanceReportUsecase: generateFinanceReportUsecase}
}

func (controller *ApiController) handleGenerateReport(context *gin.Context) {
	var financeRequest domain.FinanceReportRequest
	context.BindJSON(&financeRequest)
	context.IndentedJSON(http.StatusOK, controller.generateFinanceReportUsecase.Execute(mapper.ConvertApiRequestToCoreRequest(financeRequest)))
}

func (controller *ApiController) StartServer() {
	router := gin.Default()
	router.POST("/generateReport", controller.handleGenerateReport)
	router.Run("localhost:8080")
}
