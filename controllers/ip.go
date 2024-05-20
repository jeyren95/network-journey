package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeyren95/network-journey/utils"
	"github.com/jeyren95/network-journey/models"
	"net/http"
)

func IpHops(ctx *gin.Context) {
	var reqBody models.IpHopsReqBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
	}
	fmt.Println(reqBody)
	utils.TraceRoute(&reqBody)
}
