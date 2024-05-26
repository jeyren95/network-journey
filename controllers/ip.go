package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jeyren95/network-journey/models"
	"github.com/jeyren95/network-journey/utils"
	"net/http"
)

func IpHops(ctx *gin.Context) {
	var reqBody models.IpHopsReqBody

	// TODO: add validation for inputs
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
	}
	ipHops, err := utils.TraceRoute(&reqBody)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
	}

	geolocations, err := utils.GetGeolocations(ipHops)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ipHops":       ipHops,
		"geolocations": geolocations,
	})
}
