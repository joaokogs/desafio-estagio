package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaokogs/desafio-estagio/schemas"
)

func ListMission(ctx *gin.Context) {
	missions := []schemas.Missions{}

	if err := db.Find(&missions).Error;err!=nil{
		sendError(ctx, http.StatusInternalServerError,"Error listing missions")
		return
	}

	sendSuccess(ctx,"list-missions",missions)
}
