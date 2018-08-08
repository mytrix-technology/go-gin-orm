package main

import (
	"fmt"
	"go-gin-orm/models"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

var ORM orm.Ormer

func init() {
	models.ConnDB()
	ORM = models.GetOrmObj()
}

func main() {
	router := gin.Default()
	router.POST("/createUser", createUser)
	router.GET("/readUsers", readUsers)
	router.PUT("/updateUser", updateUser)
	router.DELETE("/deleteUser", deleteUser)
	router.Run(":3000")
}

func createUser(c *gin.Context) {
	var userBaru models.Users
	c.BindJSON(&userBaru)

	_, err := ORM.Insert(&userBaru)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"email":     userBaru.Email,
			"user_name": userBaru.UserName,
			"user_id":   userBaru.UserId})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "Gagal Input User Baru!!!"})
	}
}

func readUsers(c *gin.Context) {
	var user []models.Users
	fmt.Println(ORM)

	_, err := ORM.QueryTable("users").All(&user)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"users": &user})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "Gagal Mengambil Data User"})
	}
}

func updateUser(c *gin.Context) {
	var editUser models.Users
	c.BindJSON(&editUser)

	_, err := ORM.QueryTable("users").Filter("email", editUser.Email).Update(
		orm.Params{"user_name": editUser.UserName,
			"password": editUser.Password})

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"users": &editUser, "status": http.StatusOK})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "Gagal Update User"})
	}
}

func deleteUser(c *gin.Context) {
	var hpsUser models.Users
	c.BindJSON(&hpsUser)

	_, err := ORM.QueryTable("users").Filter("email", hpsUser.Email).Delete()

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "Gagal Hapus User"})
	}
}
