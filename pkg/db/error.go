package db

import "errors"

var (
	EntityNotFoundErr = errors.New("entity not found")
	DBErr             = errors.New("wrong database")
	ConstraintErr     = errors.New("constraint check error")
)
