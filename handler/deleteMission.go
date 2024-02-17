package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaokogs/desafio-estagio/schemas"
)

func DeleteMission(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == ""{
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id","queryParam").Error())
		return
	}

	mission := schemas.Missions{}

	if err := db.First(&mission,id).Error; err!=nil{
		sendError(ctx,http.StatusNotFound, fmt.Sprintf("Mission with id : %s not found",id))
		return
	}

	if err := db.Delete(&mission).Error;err!=nil{
		sendError(ctx,http.StatusInternalServerError,fmt.Sprintf("Error deleting mission with id : %s",id))
		return
	}

	sendSuccess(ctx,"delete-mission",mission)
}
