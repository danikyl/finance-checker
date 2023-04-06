package main

import (
	"github.com/danikyl/finance-helper/controller"
	"github.com/danikyl/finance-helper/service"
	"github.com/danikyl/finance-helper/gateway"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	
	gateway.PrintSelic()
	//injecting dependencies
	financeService := service.NewFinanceServiceImpl()
	apiController := controller.NewController(financeService)

	//running server
	apiController.StartServer()
}
