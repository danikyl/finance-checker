package controller

import (
	"net/http"

	"github.com/danikyl/finance-helper/model"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Run(host ...string) (err error)
}

func NewController() Router {
	router := gin.Default()
	router.GET("/getReport", generateReport)
	return router
}

func generateReport(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, model.FinanceReportResponse{})
}
