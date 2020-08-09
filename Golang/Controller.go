package main

import(
	"github.com/gin-gonic/gin"
	"fmt"
)

func getAllPhones(context *gin.Context)  {
	var phones []Phone
	db.Find(&phones)
	fmt.Println(phones)
	context.JSON(200,gin.H{
		"result" : "Success",
		"data" : phones,
	})

	
}

func createPhone(context *gin.Context)  {
	var phone Phone

	if err := context.ShouldBindJSON(&phone); err != nil {
		context.JSON(404,gin.H{"error":err.Error()})
		return
	}

	db.Create(&phone)
	context.JSON(200,gin.H{
		"result" : "Success",
	})

}

func updatePhone(context *gin.Context)  {

	var name = context.Param("name")
	var price = context.Param("price")
	fmt.Println(name)
	fmt.Println(price)
	db.Model(&Phone{}).Where("name = ?",name).Update("price",price)

	context.JSON(200,gin.H{
		"result" : "Success",
	})
	
}

func deletePhone(context *gin.Context)  {

	var name = context.Param("name")
	fmt.Println(name)
	db.Where("name = ?",name).Unscoped().Delete(&Phone{})

	context.JSON(200,gin.H{
		"result" : "Success",
	})
}