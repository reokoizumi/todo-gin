package models

import "time"


type Todos struct{
	Id uint 								`json:"id" gorm:"primaryKey"`
	Description string			`json:"description"`
	UpdateDate *time.Time 	`json:"updateDate"`
}