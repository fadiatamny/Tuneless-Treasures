package models

type Instrument struct {
	DBEntity
	Brand      string   `json:"brand" gorm:"column:brand"`
	Model      string   `json:"model" gorm:"column:model"`
	Price      float64  `json:"price" gorm:"column:price"`
	CategoryID string   `json:"category_id" gorm:"column:category_id"`
	Category   Category `json:"category" gorm:"foreignkey:ID"`
}
