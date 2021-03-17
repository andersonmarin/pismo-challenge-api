package model

type OperationType struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Negative bool   `json:"negative"`
}
