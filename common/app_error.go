package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnAuthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong with DB",
		err.Error(),
		"DB_ERROR",
	)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong in the server",
		err.Error(),
		"ErrInternal",
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotcreate%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot Delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot Delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}
func ErrRecordNotFound(err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("record not found"), fmt.Sprintf("RecordNotFound"))
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("invalid permission"), fmt.Sprintf("ErrNoPermission"))
}