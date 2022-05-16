package models

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/hexcraft-biz/model"
	"github.com/jmoiron/sqlx"
)

//================================================================
// Data Struct
//================================================================
type EntityScope struct {
	*model.Prototype   `dive:""`
	ResourceDomainName string `db:"resource_domain_name"`
	ResourceName       string `db:"resource_name"`
	Name               string `db:"name"`
	Type               string `db:"type"`
}

func (s *EntityScope) GetAbsScope() (*AbsScope, error) {
	return &AbsScope{
		ID:                 *s.ID,
		ResourceDomainName: s.ResourceDomainName,
		ResourceName:       s.ResourceName,
		Name:               s.Name,
		Type:               s.Type,
		CreatedAt:          s.Ctime.Format("2006-01-02 15:04:05"),
		UpdatedAt:          s.Mtime.Format("2006-01-02 15:04:05"),
	}, nil
}

type AbsScope struct {
	ID                 uuid.UUID `json:"id"`
	ResourceDomainName string    `json:"resourceDomainName"`
	ResourceName       string    `json:"resourceName"`
	Name               string    `json:"name"`
	Type               string    `json:"type"`
	CreatedAt          string    `json:"createdAt"`
	UpdatedAt          string    `json:"updatedAt"`
}

type EntityScopes []EntityScope
type AbsScopes []AbsScope

func (s *EntityScopes) GetAbsScopes() (*AbsScopes, error) {
	absScopes := AbsScopes{}

	for _, scope := range *s {
		abs, _ := scope.GetAbsScope()
		absScopes = append(absScopes, *abs)
	}

	return &absScopes, nil
}

//================================================================
// Engine
//================================================================
type ScopesTableEngine struct {
	*model.Engine
}

func NewScopesTableEngine(db *sqlx.DB) *ScopesTableEngine {
	return &ScopesTableEngine{
		Engine: model.NewEngine(db, "scopes"),
	}
}

func (e *ScopesTableEngine) Insert(resourceDomainName string, resourceName string, name string, typeStr string) (*EntityScope, error) {

	s := &EntityScope{
		Prototype:          model.NewPrototype(),
		ResourceDomainName: resourceDomainName,
		ResourceName:       resourceName,
		Name:               name,
		Type:               typeStr,
	}

	_, err := e.Engine.Insert(s)
	return s, err
}

type ScopeListQuery struct {
	ResourceDomainName string
	ResourceName       string
	Name               interface{}
	Type               string
}

func (e *ScopesTableEngine) List(listQuery ScopeListQuery, limit, offset string) (*EntityScopes, error) {
	rows := EntityScopes{}

	args := []interface{}{}
	andSqlSlice := []string{}

	if listQuery.ResourceDomainName != "" {
		andSqlSlice = append(andSqlSlice, "`resource_domain_name` = ?")
		args = append(args, listQuery.ResourceDomainName)
	}
	if listQuery.ResourceName != "" {
		andSqlSlice = append(andSqlSlice, "`resource_name` = ?")
		args = append(args, listQuery.ResourceName)
	}
	if listQuery.Type != "" {
		andSqlSlice = append(andSqlSlice, "`type` = ?")
		args = append(args, listQuery.Type)
	}

	if listQuery.Name != "" && listQuery.Name != nil {
		t := reflect.TypeOf(listQuery.Name).String()
		if t == "string" {
			andSqlSlice = append(andSqlSlice, "`name` = ?")
			args = append(args, listQuery.Name.(string))
		} else if t == "[]string" {
			orSqlSlice := []string{}
			for _, v := range listQuery.Name.([]string) {
				orSqlSlice = append(orSqlSlice, "`name` = ?")
				args = append(args, v)
			}

			orSQL := strings.Join(orSqlSlice[:], " OR ")
			andSqlSlice = append(andSqlSlice, `(`+orSQL+`)`)
		}
	}

	andSQL := "1"
	if len(andSqlSlice) >= 1 {
		andSQL = strings.Join(andSqlSlice[:], " AND ")
	}

	q := `SELECT * FROM ` + e.TblName + ` WHERE ` + andSQL + ` LIMIT ` + limit + ` OFFSET ` + offset

	if err := e.Engine.Select(&rows, q, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &rows, nil
}

func (e *ScopesTableEngine) GetByName(identity string) (*EntityScope, error) {
	row := EntityScope{}
	q := `SELECT * FROM ` + e.TblName + ` WHERE name = ?;`
	if err := e.Engine.Get(&row, q, identity); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &row, nil
}

func (e *ScopesTableEngine) DeleteByDomainName(resDomainName string) (int64, error) {
	q := `DELETE FROM ` + e.TblName + ` WHERE resource_domain_name = ?;`
	if rst, err := e.Exec(q, resDomainName); err != nil {
		return 0, err
	} else {
		return rst.RowsAffected()
	}
}
