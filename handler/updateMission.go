package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaokogs/desafio-estagio/schemas"
)

func UpdateMission(ctx *gin.Context) {
	request := UpdateMissionRequest{}

	ctx.BindJSON(&request)

	if err := request.Validade();err!=nil{
		logger.Errorf("Validation error: %v",err.Error())
		sendError(ctx,http.StatusBadRequest,err.Error())
		return
	}

	id := ctx.Query("id")

	if id == ""{
		sendError(ctx,http.StatusBadRequest, errParamIsRequired("id","queryParam").Error())
		return
	}

	mission := schemas.Missions{}

	if err :=  db.First(&mission, id).Error;err!=nil{
		sendError(ctx,http.StatusNotFound,"Mission not found")
		return
	}

	// Update Mission

	if request.Name != "" {
		mission.Name = request.Name
	}

	if request.Description != "" {
		mission.Description = request.Description
	}

	if request.Difficulty != "" {
		mission.Difficulty = request.Difficulty
	}

	if request.Reward > 0 {
		mission.Reward = request.Reward
	}

	// Save mission

	if err := db.Save(&mission).Error;err!=nil{
		logger.Errorf("Error updating mission: %v", err.Error())
		sendError(ctx,http.StatusInternalServerError, "Error updating mission")
		return
	}

	sendSuccess(ctx,"uptade-mission",mission)
	


}
