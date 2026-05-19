package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Airport struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

var airports = []Airport{
	{Code: "SGN", Name: "Tan Son Nhat International Airport", City: "Ho Chi Minh City", Country: "Viet Nam"},
	{Code: "HAN", Name: "Noi Bai International Airport", City: "Ha Noi", Country: "Viet Nam"},
}

func main() {
	r := gin.Default()

	r.GET("/airports", func(c *gin.Context) {
		c.JSON(http.StatusOK, airports)
	})

	r.GET("/airports/:code", func(c *gin.Context) {
		code := c.Param("code")
		for _, a := range airports {
			if a.Code == code {
				c.JSON(http.StatusOK, a)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Không tìm thấy sân bay!"})
	})

	r.POST("/airports", func(c *gin.Context) {
		var newAirport Airport
		if err := c.ShouldBindJSON(&newAirport); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		airports = append(airports, newAirport)
		c.JSON(http.StatusCreated, gin.H{"message": "Thêm thành công!", "data": newAirport})
	})

	r.PUT("/airports/:code", func(c *gin.Context) {
		code := c.Param("code")
		var updatedAirport Airport

		if err := c.ShouldBindJSON(&updatedAirport); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, a := range airports {
			if a.Code == code {
				airports[i] = updatedAirport
				c.JSON(http.StatusOK, gin.H{"message": "Cập nhật (PUT) thành công!", "data": airports[i]})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Không tìm thấy sân bay để cập nhật!"})
	})

	r.PATCH("/airports/:code", func(c *gin.Context) {
		code := c.Param("code")

		var partialData struct {
			Name string `json:"name"`
		}

		if err := c.ShouldBindJSON(&partialData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, a := range airports {
			if a.Code == code {
				if partialData.Name != "" {
					airports[i].Name = partialData.Name
				}
				c.JSON(http.StatusOK, gin.H{"message": "Cập nhật (PATCH) thành công!", "data": airports[i]})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Không tìm thấy sân bay!"})
	})

	r.DELETE("/airports/:code", func(c *gin.Context) {
		code := c.Param("code")
		for i, a := range airports {
			if a.Code == code {
				airports = append(airports[:i], airports[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Xóa sân bay thành công!"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Không tìm thấy sân bay để xóa!"})
	})

	r.Run(":9999")
}
