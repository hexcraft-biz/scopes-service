package features

import (
	"github.com/gin-gonic/gin"
	"github.com/hexcraft-biz/feature"
	"github.com/hexcraft-biz/scopes-service/config"
	"github.com/hexcraft-biz/scopes-service/controllers"
)

func LoadCommon(e *gin.Engine, cfg *config.Config) {
	c := controllers.NewCommon(cfg)
	e.NoRoute(c.NotFound())

	sample1v1 := feature.New(e, "/healthCheck/v1")
	sample1v1.GET("/ping", c.Ping())
}
