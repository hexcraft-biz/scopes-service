package controllers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/hexcraft-biz/controller"
	"github.com/hexcraft-biz/scopes-service/config"
	"github.com/hexcraft-biz/scopes-service/models"
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

type targetScope struct {
	Name string `uri:"name" binding:"required,min=5,max=128"`
}

func (ctrl *Scopes) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params targetScope
		if err := c.ShouldBindUri(&params); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if match := scopeNamevalidate(params.Name); match == false {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "The format of the scope name does not match."})
			return
		}

		if entityRes, err := models.NewScopesTableEngine(ctrl.DB).GetByName(params.Name); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		} else {
			if entityRes == nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
				return
			} else {
				if absRes, absErr := entityRes.GetAbsScope(); absErr != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": absErr.Error()})
					return
				} else {
					c.AbortWithStatusJSON(http.StatusOK, absRes)
					return
				}
			}
		}
	}
}

type listParams struct {
	ResourceDomainName string `form:"resourceDomainName" binding:"max=128"`
	ResourceName       string `form:"resourceName" binding:"max=128"`
	Name               string `form:"name"`
	Type               string `form:"type" binding:"oneof='' 'public' 'private'"`
	Limit              string `form:"limit,default=20"`
	Offset             string `form:"offset,default=0"`
}

func (ctrl *Scopes) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params listParams
		if err := c.ShouldBindQuery(&params); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		q := models.ScopeListQuery{
			ResourceDomainName: params.ResourceDomainName,
			ResourceName:       params.ResourceName,
			Type:               params.Type,
		}

		names := strings.Split(params.Name, "|")
		if len(names) > 1 {
			q.Name = names
		} else {
			q.Name = names[0]
		}

		if entityRes, err := models.NewScopesTableEngine(ctrl.DB).List(q, params.Limit, params.Offset); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		} else {
			if absRes, absErr := entityRes.GetAbsScopes(); absErr != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": absErr.Error()})
				return
			} else {
				c.AbortWithStatusJSON(http.StatusOK, absRes)
				return
			}
		}
	}
}

type createParams struct {
	ResourceDomainName string `json:"resourceDomainName" binding:"required,min=5,max=128"`
	ResourceName       string `json:"resourceName" binding:"required,min=5,max=128"`
	Name               string `json:"name" binding:"required,min=5,max=128"`
	Type               string `json:"type" binding:"required,oneof='public' 'private'"`
}

func (ctrl *Scopes) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params createParams
		if err := c.ShouldBindJSON(&params); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if match := scopeNamevalidate(params.Name); match == false {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "The format of the scope name does not match."})
			return
		}

		if entityRes, err := models.NewScopesTableEngine(ctrl.DB).Insert(params.ResourceDomainName, params.ResourceName, params.Name, params.Type); err != nil {
			if sqlErr, ok := err.(*mysql.MySQLError); ok && sqlErr.Number == 1062 {
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": http.StatusText(http.StatusConflict)})
				return
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		} else {
			if absRes, absErr := entityRes.GetAbsScope(); absErr != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": absErr.Error()})
				return
			} else {
				c.AbortWithStatusJSON(http.StatusOK, absRes)
				return
			}
		}
	}
}

type deleteParams struct {
	ResourceDomainName string `form:"resourceDomainName" binding:"required,min=5,max=128"`
}

func (ctrl *Scopes) DeleteByDomainName() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params deleteParams
		if err := c.ShouldBindQuery(&params); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if _, err := models.NewScopesTableEngine(ctrl.DB).DeleteByDomainName(params.ResourceDomainName); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"message": http.StatusText(http.StatusNoContent)})
			return
		}
	}
}

func scopeNamevalidate(name string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9\.\*]{0,}$`, name)
	return match
}
