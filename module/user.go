package module


type User struct {
	Email  string `form:"email" json:"email" binding:"required"` 
	Passwd string `form:"passwd" json:"passwd" binding:"required"` 
}