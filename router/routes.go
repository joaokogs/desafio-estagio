package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joaokogs/desafio-estagio/docs"
	docs "github.com/joaokogs/desafio-estagio/docs"
	"github.com/joaokogs/desafio-estagio/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(router *gin.Engine) {

	handler.InitHandler()
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
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
	// init swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
