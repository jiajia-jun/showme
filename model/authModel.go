package model

// User 用户登录项
type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

// UpdatePassword 修改密码请求
type UpdatePassword struct {
	Username    string `json:"username" binding:"required"`
	OldPassword string `json:"old_password" binding:"required,min=8"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}
