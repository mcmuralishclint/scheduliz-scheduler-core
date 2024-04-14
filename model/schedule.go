package model

import "time"

type ExecutionType string
type State string

const (
	OneTime    ExecutionType = "OneTime"
	Recurrent  ExecutionType = "Recurrent"
	Init       State         = "Init"
	InProgress State         = "InProgress"
	Done       State         = "Done"
	Failed     State         = "Failed"
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
	State         State
}
