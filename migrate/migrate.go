package main

import (
	"github.com/mtami/go-crud/config"
	"github.com/mtami/go-crud/models"
)

func init() {
	config.LoadEnvVariables()
}

func main() {

	config.GetDB().AutoMigrate(&models.Post{})

}
