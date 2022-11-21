package browser

import (
	"database/sql"
	"encoding/base64"
	"ws/pkg/types"
)

var (
	StatusError   = 0
	StatusWait    = 1
	StatusSuccess = 2
)

type Model struct {
	UUID    string `json:"uuid" db:"uuid"`
	Url     string `json:"url" db:"url"`
	Element string `json:"element" db:"element"`
	Status  int    `json:"status" db:"status"`
	Data    []byte `json:"data" db:"data"`
	types.BaseModel
}

func (m *Model) Create() (sql.Result, error) {
	return m.DB.NamedExec(`INSERT INTO screen (uuid, url, element, status, data) 
								VALUES (:uuid, :url, :element, :status, :data)`, m)
}

func (m *Model) toResource() Resource {
	return Resource{
		UUID:    m.UUID,
		Url:     m.Url,
		Element: m.Element,
		Image:   "data:image/png;base64," + base64.StdEncoding.EncodeToString(m.Data),
	}
}
