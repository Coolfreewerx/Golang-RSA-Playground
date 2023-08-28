package service

import (
	m "decode-decrypt-playground/file/model"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/joho/godotenv"
)

type RecordService struct{}

func (s *RecordService) AllRecords()([]m.OutputRecord, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Open the CSV file
	file, err := os.Open(os.Getenv("FILE_NAME"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV file into a slice of Record structs
	var records []m.Record
	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		panic(err)
	}

	output_records := []m.OutputRecord{}

	for _, record := range records {
		output_records = append(output_records, m.OutputRecord{
			Total_records:     len(records),
			No:                record.No,
			Announce_date:     record.Announce_date,
			Sex:               record.Sex,
			Age:               record.Age,
			Risk:              record.Risk,
			Province_of_onset: record.Province_of_onset,
		})
	}

	return output_records, nil
}