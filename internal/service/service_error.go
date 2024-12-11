package service

import "errors"

const ServerError = "system error"

func ServiceErrorBuilder(err error) string {
	var serviceError ServiceError
	if errors.As(err, &serviceError) {
		return err.Error()
	}
	return ServerError
}

type ServiceError error

var (
	ErrValidation ServiceError = errors.New("some of the fields not valid, revise the values")
	ErrCreation   ServiceError = errors.New("some problem occurred while inserting the item")
)
