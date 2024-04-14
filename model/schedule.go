package model

import "time"

type ExecutionType string

const (
	OneTime   ExecutionType = "OneTime"
	Recurrent ExecutionType = "Recurrent"
)

type Schedule struct {
	ID            string
	Action        string
	Schedule      string
	Description   string
	Priority      int
	Data          string
	ExecuteAt     *time.Time
	CreatedAt     time.Time
	ExecutionType ExecutionType
}
