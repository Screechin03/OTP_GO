package main

import (
	"github.com/Screechin_03/OTP_GO/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	app := api.Config{Router: router}

	app.Routes()

	router.Run(":8000")
}
