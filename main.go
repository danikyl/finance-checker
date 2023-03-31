package main

import "github.com/danikyl/finance-helper/controller"


func main() {
	router := controller.NewController()

	router.Run("localhost:8080")
}
