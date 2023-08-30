package controller

import (
	m "decode-decrypt-playground/model"
	s "decode-decrypt-playground/service"
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

// @Summary 	Show all records.
// @Description Get all records from covid-thailand-2021.csv.
// @Tags 		Gets
// @Accept 		*/*
// @Produce 	json
// @Success 	200 {array} m.AllOutputRecords "Success"
// @Failure 	500 {string} string "Internal server error"
// @Router 		/records [get]
func (c *RecordController) GetAllRecords(context echo.Context) error {
	records, err := c.recordService.GetAllRecords()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, records)
}

// @Summary 	Show record by use no.
// @Description Get a record by use no in covid-thailand-2021.csv.
// @Tags 		Gets
// @Accept 		*/*
// @Produce 	json
// @Param       no path int true "No of record"
// @Success 	200 {array} m.AllOutputRecords "Success"
// @Failure 	400 {string} string "Invalid no"
// @Failure 	500 {string} string "Internal server error"
// @Router 		/records/{no} [get]
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