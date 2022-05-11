package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hexcraft-biz/controller"
	"github.com/hexcraft-biz/scopes-service/config"
)

type Scopes struct {
	*controller.Prototype
}

func NewScopes(cfg *config.Config) *Scopes {
	return &Scopes{
		Prototype: controller.New("scopes", cfg.DB),
	}
}

func (ctrl *Scopes) NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
	}
}

func (ctrl *Scopes) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK)})
	}
}

func (ctrl *Scopes) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK)})
	}
}

func (ctrl *Scopes) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK)})
	}
}

func (ctrl *Scopes) DeleteByDomainName() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK)})
	}
}
