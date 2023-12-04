package domain

type User struct {
	ID       int    `gorm:"column:user_id"`
	Name     string `gorm:"column:name_user" json:"name"`
	Email    string `gorm:"column:mail;unique_index" json:"email"`
	Role     string `gorm:"column:rol" json:"role"`
	Password string `gorm:"column:pwd" json:"password"`
}
