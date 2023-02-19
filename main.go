package main

import (
	"github.com/mtami/go-crud/config"
	"github.com/mtami/go-crud/routers"
)

func init() {

	config.LoadEnvVariables()
	_ = config.GetDB()

}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @host      localhost:8080
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth
func main() {
	// defer config.GetDB().Close()
	api := routers.SetupRouter()

	api.Run()
}
