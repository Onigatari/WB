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
	router.LoadHTMLGlob("templates/*")

	router.GET("/", MainPage)
	router.GET("/page", GetOrderPage)
	router.GET("/order/:id", GetOrderById)

	log.Println("HTTP server is running! Connect - http://localhost:8080")
	router.Run(":8080")
}

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"database": DataBaseOrdersCache,
	})
}

func GetOrderPage(c *gin.Context) {
	id := c.Request.URL.Query()["order"][0]
	if val, ok := DataBaseOrdersCache[id]; ok {
		c.HTML(http.StatusOK, "order_page.html", gin.H{
			"order": val,
		})
	} else {
		c.IndentedJSON(http.StatusNotFound, "Error! ID does not exist")
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

func СacheFromDatabase() {
	database.CreateDbTablesIfNotExist()
	ordersFromDB := database.GetAllOrdersFromDB()

	for _, item := range ordersFromDB {
		DataBaseOrdersCache[item.OrderUID] = item
	}
}
