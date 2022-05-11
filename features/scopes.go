package features

import (
	"github.com/gin-gonic/gin"
	"github.com/hexcraft-biz/feature"
	"github.com/hexcraft-biz/scopes-service/config"
	"github.com/hexcraft-biz/scopes-service/controllers"
)

func LoadScopes(e *gin.Engine, cfg *config.Config) {
	c := controllers.NewScopes(cfg)
	e.NoRoute(c.NotFound())

	scopesV1 := feature.New(e, "/scopes/v1")

	scopesV1.GET("/scopes/:name", c.GetOne())
	scopesV1.GET("/scopes", c.List()) // skip=offset=host=&name=
	scopesV1.POST("/scopes", c.Create())
	scopesV1.DELETE("/scopes", c.DeleteByDomainName()) // resourceDomainName=
}
