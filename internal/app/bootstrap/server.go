package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/app/handler/locations"
	"github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/core"
	reposLocation "github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/infraestructure/repositories/location"
)

func StartServer() {
	e := gin.Default()
	ConfigureRoutes(e)
	err := e.Run(":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server Iniciado")

}

func ConfigureRoutes(e *gin.Engine) {
	redisClient := core.NewRedisClient("localhost", "6379", 0)
	locationRepo := reposLocation.NewLocationRepository(redisClient)
	locationHandler := locations.NewLocationHandler(locationRepo)
	g := e.Group("/api/v1")
	{
		g.GET("/states", locationHandler.FindAllStates)
		g.GET("/cities/:state", locationHandler.FindCitiesByState)
	}
}
