package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"math_api/lib"
)

type PostParams struct {
	Numbers    []float64 `json:"numbers"`
	Count      int       `json:"count"`
	Percentile int       `json:"percentile"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/min", func(c *gin.Context) {
		var data PostParams
		if err := c.BindJSON(&data); err != nil {
			return
		}

		numbers, err := lib.Min(data.Numbers, data.Count)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.IndentedJSON(http.StatusOK, numbers)
	})

	r.POST("/max", func(c *gin.Context) {
		var data PostParams
		if err := c.BindJSON(&data); err != nil {
			return
		}

		numbers, err := lib.Max(data.Numbers, data.Count)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.IndentedJSON(http.StatusOK, numbers)
	})

	r.POST("/median", func(c *gin.Context) {
		var data PostParams
		if err := c.BindJSON(&data); err != nil {
			return
		}

		median := lib.Median(data.Numbers)
		c.IndentedJSON(http.StatusOK, median)
	})

	r.POST("/avg", func(c *gin.Context) {
		var data PostParams
		if err := c.BindJSON(&data); err != nil {
			return
		}

		average := lib.Average(data.Numbers)
		c.IndentedJSON(http.StatusOK, average)
	})

	r.POST("/percentile", func(c *gin.Context) {
		var data PostParams
		if err := c.BindJSON(&data); err != nil {
			return
		}

		percentile := lib.Percentile(data.Numbers, data.Percentile)
		c.IndentedJSON(http.StatusOK, percentile)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
