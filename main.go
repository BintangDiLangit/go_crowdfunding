package main

import (
	"bwastartup/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "bintangmfhd:GHW*DT*WD(*W@tcp(108.137.5.203:3306)/golang_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("Connection established")

	// var users []user.User

	// db.Find(&users)

	// for _, user := range users {
	// 	fmt.Println(user)
	// }

	router := gin.Default()
	router.GET("/handler", handler)

	router.Run()
}

func handler(c *gin.Context) {
	dsn := "bintangmfhd:GHW*DT*WD(*W@tcp(108.137.5.203:3306)/golang_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User

	db.Find(&users)

	c.JSON(http.StatusOK, &users)
}
