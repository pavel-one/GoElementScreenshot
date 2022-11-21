package types

import "github.com/jmoiron/sqlx"

type DatabaseController struct {
	DB *sqlx.DB `json:"-" db:"-"`
}

func (c *DatabaseController) Init(db *sqlx.DB) {
	c.DB = db
}

type BaseModel struct {
	DB *sqlx.DB
}

func (m *BaseModel) Init(db *sqlx.DB) {
	m.DB = db
}
