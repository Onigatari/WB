package server

import (
	"Level0/database"
	"Level0/parser"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var DataBaseOrdersCache = make(map[string]parser.Order)

func StartServer() {
	router := gin.Default()
	router.LoadHTMLGlob("index.html")

	router.GET("/", MainPage)
	router.GET("/order/:id", GetOrderById)

	log.Println("HTTP server is running! Connect - http://localhost:8080")
	router.Run("localhost:8080")
}

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"database": DataBaseOrdersCache,
	})
}

func СacheFromDatabase() {
	database.CreateDbTablesIfNotExist()
	ordersFromDB := database.GetAllOrdersFromDB()

	for _, item := range ordersFromDB {
		DataBaseOrdersCache[item.OrderUID] = item
	}
}

func GetOrderById(c *gin.Context) {
	id := c.Param("id")
	if val, ok := DataBaseOrdersCache[id]; ok {
		c.IndentedJSON(http.StatusOK, val)
	} else {
		c.IndentedJSON(http.StatusNotFound, "Error! ID does not exist")
	}
}
