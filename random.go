package main

import (
	"fmt"
	"math/rand"
	"time"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func randomInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(max-min+1) + min)
	return rand.Intn(max-min+1) + min
}

func APIMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Server", "Golang-Gin-API")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(APIMiddleware())
	r.GET("/healthz", GetHealthCheck)
	r.GET("/sum", GetSum)

	return r
}

func GetSum(c *gin.Context) {
	a := randomInt(1, 1000)
	b := randomInt(1, 1000)

	fmt.Printf("A: %d\n", a)
	fmt.Printf("B: %d\n", b)

	fmt.Printf("A+B: %d\n", a+b)

	c.JSON(http.StatusOK, gin.H{
		"a":   a,
		"b":   b,
		"sum": a + b,
	})
}

func GetHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "up !!!"})
}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("APP_PORT")
	r := setupRouter()
	r.Run(":" + port)

}
