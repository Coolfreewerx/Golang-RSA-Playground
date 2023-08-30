package service

import (
	m "decode-decrypt-playground/model"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/joho/godotenv"
)

type RecordService struct{}
var records []m.Record

func (s *RecordService) initCSV() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Open the CSV file
	file, err := os.Open(os.Getenv("FILE_PATH"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV file into a slice of Record structs
	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		panic(err)
	}
}

func (s *RecordService) GetAllRecords()(m.AllOutputRecords, error) {
	s.initCSV()
	output_records := []m.OutputRecord{}

	for _, record := range records {
		output_records = append(output_records, m.OutputRecord {
			No: record.No,
			Announce_date: record.Announce_date,
			Sex: record.Sex,
			Age: record.Age,
			Risk: record.Risk,
			Province_of_onset: record.Province_of_onset,
			
		})
	}

	if len(output_records) == 0 {
		return m.AllOutputRecords{
			Message: "No records found",
		}, nil
	}

	all_output_records := m.AllOutputRecords{
		Message: "Get All Records Success",
		Total_record: len(output_records),
		Info_record: output_records,
	}

	return all_output_records, nil
}

func (s *RecordService) GetRecordByNo(input m.InputRecord) (m.AllOutputRecords, error) {
	s.initCSV()
	output_records := []m.OutputRecord{}

	for _, record := range records {
		if record.No == input.No {
			output_records = append(output_records, m.OutputRecord {
				No: record.No,
				Announce_date: record.Announce_date,
				Sex: record.Sex,
				Age: record.Age,
				Risk: record.Risk,
				Province_of_onset: record.Province_of_onset,
			})
		}
	}

	if len(output_records) == 0 {
		return m.AllOutputRecords{
			Message: "No records found, Please check your input",
		}, nil
	}

	all_output_records := m.AllOutputRecords{
		Message: "Get Record By No Success",
		Total_record: len(output_records),
		Info_record: output_records,
	}

	return all_output_records, nil
}