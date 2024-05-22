package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeyren95/network-journey/controllers"
)

func main() {
	r := gin.Default()
	r.POST("/ip", controllers.IpHops)

	r.Run()
}
