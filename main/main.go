package main

import (
	"../config"
	"../controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/item/:id", inDB.GetItem)
	router.GET("/items", inDB.GetItems)
	router.POST("/item", inDB.CreateItem)
	router.PUT("/item", inDB.UpdateItem)
	router.DELETE("/item/:id", inDB.DeleteItem)

	router.Run(":3000")
}
