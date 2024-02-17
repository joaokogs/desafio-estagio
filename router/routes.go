package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaokogs/desafio-estagio/handler"
)

func InitRoutes(router *gin.Engine) {

	handler.InitHandler()

	v1 := router.Group("/api/v1")
	{
		//Info da missao
		v1.GET("/mission", handler.ShowMission)

		//Cria missões
		v1.POST("/mission", handler.CreateMission)

		//Deleta uma missão
		v1.DELETE("/mission", handler.DeleteMission)

		//Edita uma missão
		v1.PUT("/mission", handler.UpdateMission)

		//Lista todas as missões
		v1.GET("/missions", handler.ListMission)

	}
}
