package request

type CreateUser struct {
	Username string `json:"username,omitempty" binding:"required" form:"username"`
	Password string `json:"password,omitempty" binding:"required" form:"password"`
	FullName string `json:"full_name,omitempty" binding:"required" form:"full_name"`
	Email    string `json:"email,omitempty" binding:"required" form:"email"`
}

type Login struct {
	Username string `json:"username,omitempty" binding:"required" form:"username"`
	Password string `json:"password,omitempty" binding:"required" form:"password"`
}
