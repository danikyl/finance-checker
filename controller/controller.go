package controller

import (
	"net/http"

	"github.com/danikyl/finance-helper/model"
	"github.com/danikyl/finance-helper/service"
	"github.com/gin-gonic/gin"
)

type ApiController struct {
	financeService service.FinanceService
}

func NewController(financeServer service.FinanceService) *ApiController {
	return &ApiController{financeService: financeServer}
}

func (c *ApiController) handleGenerateReport(context *gin.Context) {
	var financeRequest model.FinanceReportRequest
	context.BindJSON(&financeRequest)
	context.IndentedJSON(http.StatusOK, c.financeService.GenerateReport(financeRequest))
}

func (c *ApiController) StartServer() {
	router := gin.Default()
	router.POST("/generateReport", c.handleGenerateReport)
	router.Run("localhost:8080")
}
