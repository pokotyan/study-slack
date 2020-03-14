package models

import "time"

type SettingHistory struct {
	ID          int       `json:"id" gorm:"primary_key"`
	SearchRange int       `json:"search_range"`
	NumOfPeople int       `json:"num_of_people"`
	Word        string    `json:"word"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
