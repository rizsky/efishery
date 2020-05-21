package routes

import (
	"github.com/allegro/bigcache"
	"github.com/efishery/fetchin-service/module/fetch/handlers"
	"github.com/efishery/fetchin-service/module/fetch/helpers"
	"github.com/efishery/fetchin-service/module/fetch/middlewares"
	"github.com/efishery/fetchin-service/module/fetch/repositories"
	"github.com/efishery/fetchin-service/module/fetch/services"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//SetupRouter untuk setup routernya bro
func SetupRouter(cache *bigcache.BigCache) *gin.Engine {

	cacheRepo := repositories.NewCacheRepo(cache)
	efisheryRepo := repositories.NewEfisheryCallRepository(helpers.EfisheryURL)
	converterRepo := repositories.NewConverterRepository(helpers.ConverterURL)
	agregatorServices := services.NewAgregatorServices(efisheryRepo, converterRepo, cacheRepo)
	handler := handlers.NewHandler(agregatorServices)
	auth := middlewares.AuthenticationRequired()
	admin := middlewares.AuthenticationRequired("admin")

	r := gin.Default()
	v1 := r.Group("/v1")
	v1.Use(auth)
	{
		v1.GET("private-claim", handler.GetPrivateClaim)
		v1.GET("list", handler.GetList)

	}
	v1Admin := r.Group("/v1/admin")
	v1Admin.Use(admin)
	{
		v1Admin.GET("list", handler.GetByParam)
	}

	r.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
