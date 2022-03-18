package domain

import (
	"errors"
)

type CustomerResponse struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}

func NewError(errMessage string) error {
	return errors.New(errMessage)
}
