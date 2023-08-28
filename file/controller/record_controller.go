package controller

import (
	m "decode-decrypt-playground/file/model"
	s "decode-decrypt-playground/file/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RecordController struct{
	recordService s.RecordService
}

func NewRecordController() *RecordController {
	return &RecordController{}
}

func (c *RecordController) GetAllRecords(context echo.Context) error {
	records, err := c.recordService.GetAllRecords()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, records)
}

func (c *RecordController) GetRecordByNo(context echo.Context) error {
	no, err := strconv.Atoi(context.Param("no"))
	if err != nil {
		return context.JSON(http.StatusBadRequest, m.AllOutputRecords{
			Message: "Invalid no",
		})
	}

	input := m.InputRecord{
		No: no,
	}

	records, err := c.recordService.GetRecordByNo(input)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, records)

}