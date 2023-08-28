package controller

import (

	s "decode-decrypt-playground/file/service"
	"net/http"


	"github.com/labstack/echo/v4"
)

type RecordController struct{
	recordService s.RecordService
}

func NewRecordController() *RecordController {
	return &RecordController{}
}

func (c *RecordController) GetAllRecords(context echo.Context) error {
	records, err := c.recordService.AllRecords()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, records)
}


	// var records_bangkok []model.Record
 
	// Print the records
	// for _, record := range records {
	// 	if record.Province == "กรุงเทพมหานคร" {
	// 		fmt.Printf("Name: %s, Sex: %s, Age: %d, Province: %s\n", record.Announce_date, record.Sex, record.Age, record.Province)
	// 		records_bangkok = append(records_bangkok, record)
	// 	}
	// }

	// total_of_woman, total_of_men := 0, 0

	// for _, record_bankok := range records_bangkok {
	// 	if record_bankok.Sex == "หญิง" {
	// 		total_of_woman += 1
	// 	} else if record_bankok.Sex == "ชาย" {
	// 		total_of_men += 1
	// 	}
	// }
	// fmt.Println()
	// fmt.Println("จำนวนผู้ติดเชื้อในกรุงเทพมหานคร: ", len(records_bangkok), "ราย")
	// fmt.Println("เป็นผู้หญิง: ", total_of_woman, "ราย", " เป็นผู้ชาย: ", total_of_men)