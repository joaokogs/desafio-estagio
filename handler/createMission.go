package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaokogs/desafio-estagio/schemas"
)

func CreateMission(ctx *gin.Context) {

	request := CreateMissionRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validade error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	mission := schemas.Missions{

		Name : request.Name,
		Description: request.Description,
		Difficulty: request.Difficulty,
		Reward: request.Reward,
	}

	if err := db.Create(&mission).Error; err != nil {
		logger.Errorf("error creating mission: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError,"error creating missions on database")
		return
	}

	sendSuccess(ctx,"CreateMission",mission)
}
