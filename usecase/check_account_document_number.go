package usecase

import (
	"errors"
	"regexp"
)

const AccountDocumentNumberLength = 11

var (
	ErrAccountDocumentNumberNotAccepted = errors.New("document number not accepted")
	AccountDocumentNumberFormat         = regexp.MustCompile("\\D")
)

func CheckAccountDocumentNumber(documentNumber string) error {
	if len(documentNumber) != AccountDocumentNumberLength {
		return ErrAccountDocumentNumberAlreadyExists
	}

	if AccountDocumentNumberFormat.MatchString(documentNumber) {
		return ErrAccountDocumentNumberNotAccepted
	}

	return nil
}
