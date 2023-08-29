package model

type Record struct { 
	No 						int 	`csv:"No."`
	Announce_date 			string 	`csv:"announce_date"`
	Sex 					string 	`csv:"sex"` 
	Age 					int 	`csv:"age"`
	Unit 					string 	`csv:"unit"`
	Province_of_isolation 	string 	`csv:"province_of_isolation"`
	Risk 					string 	`csv:"risk"`
	Nationality     		string	`csv:"nationality"`
	Province_of_onset		string  `csv:"province_of_onset"`
	District_of_onset		string  `csv:"district_of_onset"`
}

type InputRecord struct {
	No 				int 	`json:"no"`
	Province	    string  `json:"province"`
	Age				string  `json:"age"`
}

type OutputRecord struct {
	No 						int 	`json:"no"`
	Announce_date 			string 	`json:"announce_date"`
	Sex						string  `json:"sex"`
	Age 					int 	`json:"age"`
	Risk 					string 	`json:"risk"`
	Province_of_onset		string  `json:"province_of_onset"`
}

type AllOutputRecords struct {
	Message 			string 			`json:"message"`
	Total_record 		int  			`json:"total_record"`
	Info_record		    []OutputRecord 	`json:"info_record"`
}