package controllers
import (
 models "tableFootballGo/models"
 "net/http"
"github.com/gin-gonic/gin"
 "github.com/jinzhu/gorm"
)
// GET /users
// Get all users
func FindUsers(c *gin.Context) {
 db := c.MustGet("db").(*gorm.DB)
var users []models.User
 db.Find(&users)
c.JSON(http.StatusOK, gin.H{"data": users})
}
// POST /users
// Create new users
func CreateUser(c *gin.Context) {
 db := c.MustGet("db").(*gorm.DB)
// Validate input
 var input models.CreateUserInput
 if err := c.ShouldBindJSON(&input); err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  return
 }
// Create User
 user := models.User{Name: input.Name, Surname: input.Surname}
 db.Create(&user)
c.JSON(http.StatusOK, gin.H{"data": user})
}
// GET /users/:id
// Find a users
func FindUser(c *gin.Context) {
 db := c.MustGet("db").(*gorm.DB)
// Get model if exist
 var user models.User
 if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
  return
 }
c.JSON(http.StatusOK, gin.H{"data": user})
}
// PATCH /users/:id
// Update a user
func UpdateUser(c *gin.Context) {
 db := c.MustGet("db").(*gorm.DB)
// Get model if exist
 var user models.User
 if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
  return
 }
// Validate input
 var input models.UpdateUserInput
 if err := c.ShouldBindJSON(&input); err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  return
 }
db.Model(&user).Updates(input)
c.JSON(http.StatusOK, gin.H{"data": user})
}
// DELETE /users/:id
// Delete a user
func DeleteUser(c *gin.Context) {
 db := c.MustGet("db").(*gorm.DB)
// Get model if exist
 var user models.User
 if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
  return
 }
db.Delete(&user)
c.JSON(http.StatusOK, gin.H{"data": true})
}