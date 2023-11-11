package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hexcraft-biz/scopes-service/config"
	"github.com/hexcraft-biz/scopes-service/features"
)

func main() {
	cfg, err := config.Load()
	MustNot(err)

	if cfg.GinMode == "debug" && cfg.AutoCreateDBSchema {
		MustNot(cfg.Mysql.DBInit("./sql/", []string{
			"v1.0__schema.sql",
		}))
	}

	MustNot(cfg.Mysql.Open())
	defer cfg.Mysql.Close()

	GetEngine(cfg).Run(":" + cfg.AppPort)
}

func GetEngine(cfg *config.Config) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{cfg.TrustProxy})

	features.LoadCommon(r, cfg)
	features.LoadScopes(r, cfg)
	return r
}

func MustNot(err error) {
	if err != nil {
		panic(err.Error())
	}
}
