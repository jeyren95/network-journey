package main

import (
	"github.com/jeyren95/network-journey/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/ip", controllers.IpHops)
	
	r.Run()
}
