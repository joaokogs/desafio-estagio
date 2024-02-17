package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaokogs/desafio-estagio/schemas"
)

func ShowMission(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == ""{
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id","queryParam").Error())
		return
	}

	mission := schemas.Missions{}

	if err := db.First(&mission,id).Error;err!=nil{
		sendError(ctx, http.StatusNotFound, "Mission not found")
		return
	}

	sendSuccess(ctx, "show-mission",mission)
}
