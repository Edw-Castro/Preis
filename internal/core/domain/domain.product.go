package domain

type Product struct {
	Article_id int    `gorm:"column:article_id" json:"id"`
	Name       string `gorm:"column:name_article" json:"name"`
	Brand      string `gorm:"column:brand" json:"brand"`
	// Price           float32 `gorm:"column:price" json:"price"`
	Category string `gorm:"column:category" json:"category"`
	// Content         float32 `gorm:"column:content" json:"content"`
	// UnitMeasurement string  `gorm:"column:unitMeasurement" json:"unit_measurement"`
}
