package models

import (
	"crypto/rand"
	"database/sql"
	"io"

	"github.com/google/uuid"
	"github.com/hexcraft-biz/model"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

const (
	PW_SALT_BYTES = 16
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

func (u *EntityScope) GetAbsScope() (*AbsScope, error) {
	return &AbsScope{
		ID:                 *u.ID,
		ResourceDomainName: u.ResourceDomainName,
		ResourceName:       u.ResourceName,
		Name:               u.Name,
		Type:               u.Type,
		CreatedAt:          u.Ctime.Format("2006-01-02 15:04:05"),
		UpdatedAt:          u.Mtime.Format("2006-01-02 15:04:05"),
	}, nil
}

type AbsScope struct {
	ID                 uuid.UUID `json:"id"`
	ResourceDomainName string    `json:"resource_domain_name"`
	ResourceName       string    `json:"resource_name"`
	Name               string    `json:"name"`
	Type               string    `json:"type"`
	CreatedAt          string    `json:"created_at"`
	UpdatedAt          string    `json:"updated_at"`
}

//================================================================
// Engine
//================================================================
type ScopesTableEngine struct {
	*model.Engine
}

func NewUsersTableEngine(db *sqlx.DB) *ScopesTableEngine {
	return &ScopesTableEngine{
		Engine: model.NewEngine(db, "users"),
	}
}

func (e *ScopesTableEngine) Insert(identity string, password string, status string) (*EntityScope, error) {
	saltBytes := make([]byte, PW_SALT_BYTES)
	if _, err := io.ReadFull(rand.Reader, saltBytes); err != nil {
		return nil, err
	}
	salt := string(saltBytes)

	pwdBytes := []byte(password + salt)

	hashBytes, hashErr := bcrypt.GenerateFromPassword(pwdBytes, bcrypt.DefaultCost)
	if hashErr != nil {
		return nil, hashErr
	}

	u := &EntityScope{
		Prototype: model.NewPrototype(),
		Identity:  identity,
		Password:  hashBytes,
		Salt:      saltBytes,
		Status:    status,
	}

	_, err := e.Engine.Insert(u)
	return u, err
}

func (e *ScopesTableEngine) List(id string) (*EntityScope, error) {
	rows := []EntityScope{}
	q := `SELECT * FROM ` + e.TblName + ` WHERE id = UUID_TO_BIN(?);`
	if err := e.Engine.Get(&row, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &row, nil
}

func (e *ScopesTableEngine) GetByID(id string) (*EntityScope, error) {
	row := EntityScope{}
	q := `SELECT * FROM ` + e.TblName + ` WHERE id = UUID_TO_BIN(?);`
	if err := e.Engine.Get(&row, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &row, nil
}

func (e *ScopesTableEngine) GetByIdentity(identity string) (*EntityScope, error) {
	row := EntityScope{}
	q := `SELECT * FROM ` + e.TblName + ` WHERE identity = ?;`
	if err := e.Engine.Get(&row, q, identity); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &row, nil
}

func (e *ScopesTableEngine) ResetPwd(id *uuid.UUID, password string, saltBytes []byte) (int64, error) {
	salt := string(saltBytes)

	pwdBytes := []byte(password + salt)

	hashBytes, hashErr := bcrypt.GenerateFromPassword(pwdBytes, bcrypt.DefaultCost)
	if hashErr != nil {
		return 0, hashErr
	}

	q := `UPDATE ` + e.TblName + ` SET password = ? WHERE id = UUID_TO_BIN(?);`
	if rst, err := e.Exec(q, hashBytes, &id); err != nil {
		return 0, err
	} else {
		return rst.RowsAffected()
	}
}

func (e *ScopesTableEngine) UpdateStatus(id *uuid.UUID, status string) (int64, error) {
	q := `UPDATE ` + e.TblName + ` SET status = ? WHERE id = UUID_TO_BIN(?);`
	if rst, err := e.Exec(q, status, &id); err != nil {
		return 0, err
	} else {
		return rst.RowsAffected()
	}
}
