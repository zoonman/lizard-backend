package main

import (
	"github.com/gin-gonic/gin"
	//validator "github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "github.com/zoonman/lizard-backend/api"
	. "github.com/zoonman/lizard-backend/models"
)

// var Db *gorm.Db

func main() {

	var err error
	Db, err = gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/bmi_tracker?charset=utf8&parseTime=True")
	if err == nil {
		defer Db.Close()
		Db.LogMode(true)
		Db.AutoMigrate(&User{})
		Db.AutoMigrate(&Record{})

		router := gin.Default()
		apiGroupV1 := router.Group("api/v1")
		{
			apiGroupV1.GET("/users/:id", GetUser)
			apiGroupV1.GET("/users/:id/records", GetUserRecords)
			apiGroupV1.GET("/users", GetUsers)
			apiGroupV1.POST("/users", PostUser)
			apiGroupV1.POST("/users/:id/records", PostUserRecords)
		}
		router.Run(GetEnvVar("LIZARD_ADDR", ":10123"))
	} else {
		panic(err)
	}

}
