package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func main() {
	r := gin.Default()

	r.GET("/add/:a/:b", func(c *gin.Context) {
		a, _ := strconv.Atoi(c.Param("a"))
		b, _ := strconv.Atoi(c.Param("b"))
		result := Add(a, b)
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	r.GET("/subtract/:a/:b", func(c *gin.Context) {
		a, _ := strconv.Atoi(c.Param("a"))
		b, _ := strconv.Atoi(c.Param("b"))
		result := Subtract(a, b)
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	r.GET("/multiply/:a/:b", func(c *gin.Context) {
		a, _ := strconv.Atoi(c.Param("a"))
		b, _ := strconv.Atoi(c.Param("b"))
		result := Multiply(a, b)
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	r.Run(":8181")
}
