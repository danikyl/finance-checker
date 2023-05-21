package main

import (
	"github.com/danikyl/finance-helper/api/controller"
	"github.com/danikyl/finance-helper/core/usecase"
	"github.com/danikyl/finance-helper/gateway/integration"
)

func main() {
	
	integration.PrintSelic()
	//injecting dependencies
	generateFinanceReportUsecase := usecase.NewGenerateFinanceReportImpl()
	apiController := controller.NewController(generateFinanceReportUsecase)

	//running server
	apiController.StartServer()
}
