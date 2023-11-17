package domain

type User struct {
	ID       int    `gorm:"column:id"`
	Name     string `gorm:"column:name" json:"name"`
	LastName string `gorm:"column:lastName" json:"lastName"`
	Email    string `gorm:"column:email;unique_index" json:"email"`
	Role     string `gorm:"column:role" json:"role"`
	Password string `gorm:"column:password" json:"password"`
}
