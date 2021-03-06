package handler

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/calmato/gran/api/user/internal/application/response"
	"github.com/calmato/gran/api/user/internal/domain"
)

// ErrorHandling - エラーレスポンスを返す
func ErrorHandling(ctx *gin.Context, err error) {
	res := errorResponse(err)

	ctx.JSON(res.StatusCode, res)
	ctx.Abort()
}

// errorResponse - エラー用のレスポンスを返す
func errorResponse(err error) *response.ErrorResponse {
	var res *response.ErrorResponse

	switch getErrorCode(err) {
	// 400
	case domain.UnableParseJSON:
		res = response.BadRequest
		logging("info", "Unable parse request body", err)
	case domain.NotEqualRequestWithDatastore:
		res = response.BadRequest
		logging("info", "Not Equal Request with Datastore", err)
	case domain.InvalidRequestValidation:
		res = response.BadRequest
		setValidationErrors(res, err)
		logging("info", "Invalid request validation", err, res.ValidationErrors...)
	case domain.InvalidDomainValidation:
		res = response.BadRequest
		setValidationErrors(res, err)
		logging("info", "Invalid domain validation", err, res.ValidationErrors...)
	// 401
	case domain.Unauthorized:
		res = response.Unauthorized
		logging("info", "Unauthorized", err)
	// 403
	case domain.Forbidden:
		res = response.Forbidden
		logging("info", "Forbidden", err)
	// 404
	case domain.NotFound:
		res = response.NotFound
		logging("info", "Not found", err)
	// 409
	case domain.AlreadyExists:
		res = response.AlreadyExists
		setValidationErrors(res, err)
		logging("info", "Already exists request", err, res.ValidationErrors...)
	// 500
	case domain.ErrorInDatastore:
		res = response.InternalServerError
		logging("warning", "Error in datastore", err)
	case domain.ErrorInStorage:
		res = response.InternalServerError
		logging("warning", "Error in storage", err)
	default:
		res = response.InternalServerError
		logging("error", "Internal server error", err)
	}

	res.ErrorCode = getErrorCode(err)
	return res
}

func logging(level string, message string, err error, ves ...*response.ValidationError) {
	log.Printf("%s: %s: %s", level, message, err.Error())

	if len(ves) > 0 {
		j, _ := json.Marshal(ves)

		log.Printf("%s: %s", level, j)
	}
}

func getErrorCode(err error) domain.ErrorCode {
	if e, ok := err.(domain.ShowError); ok {
		return e.Code()
	}

	return domain.Unknown
}

func setValidationErrors(er *response.ErrorResponse, err error) {
	if e, ok := err.(domain.ShowError); ok {
		ves := e.Validation()
		er.ValidationErrors = make([]*response.ValidationError, len(ves))

		for i, ve := range ves {
			er.ValidationErrors[i] = &response.ValidationError{
				Field:   ve.Field,
				Message: ve.Message,
			}
		}
	}
}
