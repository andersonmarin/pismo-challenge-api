package model

type Account struct {
	ID             uint64 `gorm:"primaryKey" json:"id"`
	DocumentNumber string `gorm:"size:11;uniqueIndex" json:"document_number"`
}
