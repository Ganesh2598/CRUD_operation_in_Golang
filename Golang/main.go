package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *gorm.DB = nil
var err error

func main()  {
	

	db, err = gorm.Open("mysql","root:kratos@/some")
	if err != nil {
		log.Fatal("Not Connected",err)
	}
	defer db.Close()
	db.AutoMigrate(&Phone{})
	db.LogMode(true)

	router := gin.Default()
	router.GET("/getphones",getAllPhones)
	router.POST("/createphone",createPhone)
	router.PATCH("/updatephone/:name/:price",updatePhone)
	router.DELETE("/deletephone/:name",deletePhone)
	  
	router.Run(":8000")
}


