package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
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
	return r
}

func TestAdd(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/add/1/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"result":2}`, w.Body.String())
}

func TestSubtract(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/subtract/2/2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"result":0}`, w.Body.String())
}

func TestMultiply(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/multiply/3/3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"result":9}`, w.Body.String())
}
