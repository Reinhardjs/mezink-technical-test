package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/reinhardjs/mezink/domain"
	"github.com/reinhardjs/mezink/domain/dto/request"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// ResponseSuccess represent the response success struct
type ResponseSuccess struct {
	Code    int             `json:"code"`
	Msg     string          `json:"msg"`
	Records []domain.Record `json:"records"`
}

// RecordHandler  represent the httphandler for record
type RecordHandler struct {
	RUsecase domain.RecordUsecase
}

// NewRecordHandler will initialize the records/resources endpoint
func NewRecordHandler(e *echo.Echo, us domain.RecordUsecase) {
	handler := &RecordHandler{
		RUsecase: us,
	}
	e.POST("/records", handler.GetBySumRange)
}

// GetBySumRange will get record by filter request
func (a *RecordHandler) GetBySumRange(c echo.Context) error {
	ctx := c.Request().Context()

	var recordRequest request.GetRecordRequest

	err := c.Bind(&recordRequest)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{
			Code: getStatusCode(err), Msg: err.Error(),
		})
	}

	records, err := a.RUsecase.GetBySumRange(ctx, &recordRequest)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{
			Code: getStatusCode(err), Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponseSuccess{
		Code:    0,
		Msg:     "Success",
		Records: records,
	})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
