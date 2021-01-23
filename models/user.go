package models
type User struct {
 ID     uint   `json:"id" gorm:"primary_key"`
 Name  string `json:"name"`
 Surname string `json:"surname"`
}
type CreateUserInput struct {
 Name  string `json:"name" binding:"required"`
 Surname string `json:"surname" binding:"required"`
}
type UpdateUserInput struct {
 Name  string `json:"name"`
 Surname string `json:"surname"`
}
