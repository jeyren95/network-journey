package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jeyren95/network-journey/controllers"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/ip", controllers.IpHops)

	r.Run()
}
