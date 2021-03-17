package model

import (
	"time"
)

type Transaction struct {
	ID              uint64    `gorm:"primaryKey" json:"id"`
	AccountID       uint64    `json:"account_id"`
	OperationTypeID uint64    `json:"operation_type_id"`
	Amount          float64   `gorm:"type:decimal(10, 2)" json:"amount"`
	EventDate       time.Time `json:"event_date"`

	Account       *Account       `json:"account,omitempty"`
	OperationType *OperationType `json:"operation_type,omitempty"`
}
