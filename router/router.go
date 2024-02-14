package router

import "github.com/gin-gonic/gin"

func Init() {
	//iniciando Router

	router := gin.Default()

	//iniciando as rotas
	InitRoutes(router)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
