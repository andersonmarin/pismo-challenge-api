package usecase

import (
	"errors"
	"regexp"
)

var ErrAccountDocumentNumberNotAccepted = errors.New("document number not accepted")

func CheckAccountDocumentNumber(documentNumber string) error {
	r, err := regexp.Compile("\\D")
	if err != nil {
		return err
	}

	if r.MatchString(documentNumber) {
		return ErrAccountDocumentNumberNotAccepted
	}

	return nil
}
