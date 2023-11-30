package domain

type Store struct {
	Store_id   string `gorm:"column:store"`
	Name_store string `gorm:"column:name_store" json:"name_store"`
	Address    string `gorm:"column:address" json:"address"`
	User_id    string `gorm:"column:user_id"`
}
