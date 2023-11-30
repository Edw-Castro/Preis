package domain

type StoreArticlePrice struct {
	Id_price        int     `gorm:"id_price"`
	Store_id        int     `gorm:"store_id"`
	Article_id      int     `gorm:"article_id"`
	Price           float32 `gorm:"price"`
	Content         int     `gorm:"content"`
	UnitMeasurement string  `gorm:"unit_measurement"`
}
