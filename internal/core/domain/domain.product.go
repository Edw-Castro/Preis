package domain

type Product struct {
	ID              int     `gorm:"column:id" json:"id"`
	Name            string  `gorm:"column:name" json:"name"`
	Brand           string  `gorm:"column:brand" json:"brand"`
	Price           float32 `gorm:"column:price" json:"price"`
	Category        string  `gorm:"column:category" json:"category"`
	Content         float32 `gorm:"column:content" json:"content"`
	UnitMeasurement string  `gorm:"column:unitmeasurement" json:"unit_measurement"`
}
