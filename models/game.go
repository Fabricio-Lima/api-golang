package models

type Game struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	Release_Date string `json:"release_date"`
	Publisher    string `json:"publisher"`
}
