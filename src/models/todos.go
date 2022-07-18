package models

import "time"


type Todos struct{
	Id uint 								`json:"id"`
	Description string			`json:"description"`
	UpdateDate *time.Time 	`json:"updateDate"`
}