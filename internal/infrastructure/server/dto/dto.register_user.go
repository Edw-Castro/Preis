package dto

type ResponseUserGet struct {
	Name string `form:"name" binding:"required"`
	// LastName string `form:"lastname" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func BuildResponseUserGet(user ResponseUserGet) ResponseUserGet {
	return ResponseUserGet(user)
}
