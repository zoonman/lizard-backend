package api

import "github.com/gin-gonic/gin"
import (
	"github.com/jinzhu/gorm"
	. "github.com/zoonman/lizard-backend/models"
)

var Db *gorm.DB

func PostUser(c *gin.Context) {
	var user User
	var err = c.BindJSON(&user)
	if err == nil {
		Db.Save(&user)
		c.JSON(200, user)
	} else {
		c.JSON(503, err)
	}
}

func GetUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	Db.First(&user, id)
	if user.ID != 0 {
		c.JSON(200, PublicUser{User: &user})
	} else {
		c.JSON(404, gin.H{"error": "User not found."})
	}
}

func GetUserRecords(c *gin.Context) {
	var user User
	var records []Record
	id := c.Params.ByName("id")
	Db.First(&user, id).Related(&records)
	if user.ID != 0 {
		c.JSON(200, records)
	} else {
		c.JSON(404, gin.H{"error": "User not found."})
	}
}

func PostUserRecords(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	Db.First(&user, id)
	if user.ID != 0 {
		var record Record
		var err = c.BindJSON(&record)
		if err == nil {
			record.UserID = user.ID
			Db.Save(&record)
			c.JSON(200, record)
		} else {
			c.JSON(503, err)
		}
	} else {
		c.JSON(404, gin.H{"error": "User not found."})
	}
}

func GetUsers(c *gin.Context) {
	var users []User
	Db.Limit(10).Find(&users)
	c.JSON(200, &users)
}
