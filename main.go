package main

import (
	"github.com/danikyl/finance-helper/controller"
	"github.com/danikyl/finance-helper/service"
)

func main() {
	//injecting dependencies
	financeService := service.NewFinanceServiceImpl()
	apiController := controller.NewController(financeService)

	//running server
	apiController.StartServer()
}
