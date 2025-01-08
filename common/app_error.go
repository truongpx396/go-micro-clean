package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// {
// 	"error": {
// 	  "code": "RESOURCE_NOT_FOUND",
// 	  "message": "The requested resource was not found.",
// 	  "details": [
// 		{
// 		  "field": "id",
// 		  "issue": "Invalid or missing"
// 		}
// 	  ],
// 	  "status": 404
// 	}
//   }

var caser = cases.Title(language.English)

type AppError struct {
	Status  int    `json:"status"`
	RootErr error  `json:"-"`
	Message string `json:"message"`
	Details string `json:"details"`
	Code    string `json:"code"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		Status:  http.StatusBadRequest,
		RootErr: root,
		Message: msg,
		Details: log,
		Code:    key,
	}
}

func WriteErrorResponse(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.Status, appErr)
		return
	}

	c.JSON(http.StatusInternalServerError, ErrInternal(err))
}

func NewBadRequestResponse(root error, msg, details, code string) *AppError {
	return &AppError{
		Status:  http.StatusBadRequest,
		RootErr: root,
		Message: msg,
		Details: details,
		Code:    code,
	}
}

func NewGeneralErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		Status:  statusCode,
		RootErr: root,
		Message: msg,
		Details: log,
		Code:    key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		Status:  http.StatusUnauthorized,
		RootErr: root,
		Message: msg,
		Code:    key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func NewBadRequestError(root error, msg string, code string) *AppError {
	if root != nil {
		return NewBadRequestResponse(root, msg, root.Error(), code)
	}

	return NewBadRequestResponse(errors.New(msg), msg, msg, code)
}

func NewInternalServerError(root error, msg string, code string) *AppError {
	if root != nil {
		return NewGeneralErrorResponse(http.StatusInternalServerError, root, msg, root.Error(), code)
	}

	return NewGeneralErrorResponse(http.StatusInternalServerError, errors.New(msg), msg, msg, code)
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
	return NewBadRequestResponse(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewBadRequestResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInvalidRequestWithMsg(err error, msg string) *AppError {
	return NewBadRequestResponse(err, msg, err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewGeneralErrorResponse(http.StatusInternalServerError, err,
		"something went wrong in the server", err.Error(), "ErrInternal")
}

func ErrInternalWithMsg(err error, msg string) *AppError {
	return NewGeneralErrorResponse(http.StatusInternalServerError, err,
		msg, err.Error(), "ErrInternal")
}

func ErrUnauthorized(err error) *AppError {
	return NewUnauthorized(err, "The request could not be authorized", "ErrUnauthorized")
}

func ErrNoPermission(err error) *AppError {
	return NewBadRequestError(
		err,
		fmt.Sprintf("You have no permission"),
		fmt.Sprintf("ErrNoPermission"),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewInternalServerError(
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", caser.String(entity)),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewBadRequestError(
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("Err%sAlreadyExists", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewBadRequestError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewGeneralErrorResponse(
		http.StatusNotFound,
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		err.Error(),
		fmt.Sprintf("Err%sNotFound", caser.String(entity)),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewInternalServerError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", caser.String(entity)),
	)
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewBadRequestError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewInternalServerError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", caser.String(entity)),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewBadRequestError(
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}

// ErrRecordNotFound is used to make our application logic independent of other libraries errors
var ErrRecordNotFound = errors.New("record not found")
