package models

import "time"

type Cell struct {
	Mine    bool `json:"mine"`
	Flaged  bool `json:"flaged"`
	Clicked bool `json:"clicked"`
	CellPoint
}

type CellPoint struct {
	RowValue    int `json:"rowValue"`
	ColumnValue int `json:"columnValue"`
}

type CellGrid []Cell

type Status string

const (
	STARTED Status = "STARTED"
	PAUSED  Status = "PAUSED"
	DEFEAT  Status = "DEFEAT"
	WON     Status = "WON"
)

type Game struct {
	GameId    string     `json:"gameId" sql:",pk"`
	PlayerId  string     `json:"playerId"`
	Rows      int        `json:"rows"`
	Cols      int        `json:"cols"`
	Mines     int        `json:"mines"`
	Status    Status     `json:"status"`
	Grid      []CellGrid `json:"grid,omitempty"`
	Clicks    int        `json:"-"`
	Timer     time.Time     `json:"timer"`
	TableName struct{}   `sql:"game_mine"`
}

type Player struct {
	Id        string `json:"Id" sql:",pk"`
	UserName  string
	Games     []*Game
	TableName struct{} `sql:"player"`
}
