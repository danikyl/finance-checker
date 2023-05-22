package controller

import (
	"net/http"

	"github.com/danikyl/finance-helper/api/domain"
	"github.com/danikyl/finance-helper/api/mapper"
	"github.com/danikyl/finance-helper/core/usecase"
	"github.com/gin-gonic/gin"
)

type apiControllerImpl struct {
	generateFinanceReportUsecase usecase.GenerateFinanceReport
}

func NewController(generateFinanceReportUsecase usecase.GenerateFinanceReport) *apiControllerImpl {
	return &apiControllerImpl{generateFinanceReportUsecase: generateFinanceReportUsecase}
}

func (controller *apiControllerImpl) handleGenerateReport(context *gin.Context) {
	var financeRequest domain.FinanceReportRequest
	context.BindJSON(&financeRequest)
	var financeResponse = controller.generateFinanceReportUsecase.Execute(mapper.ConvertApiRequestToCoreRequest(financeRequest))
	context.IndentedJSON(http.StatusOK, mapper.ConvertCoreResponseToApiResponse(*financeResponse))
}

func (controller *apiControllerImpl) StartServer() {
	router := gin.Default()
	router.POST("/generateReport", controller.handleGenerateReport)
	router.Run("localhost:8080")
}
