package controller

import (
	"fmt"
	"net/http"

	"github.com/beruang43221/assignment2-015/database"
	"github.com/beruang43221/assignment2-015/models"
	"github.com/gin-gonic/gin"
)

// GetAll Data
func GetAllData(c *gin.Context) {
	var data []models.Order
	db := database.GetDB()

	if err := db.Preload("Items").Find(&data).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, data)
}

// Create Data
func CreateData(c *gin.Context) {
	var inputData models.Order
	db := database.GetDB()

	if err := c.ShouldBindJSON(&inputData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	if err := db.Create(&inputData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}

	c.JSON(http.StatusCreated, inputData)
}

// Update Data {PUT}
func UpdateData(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
 
	var order models.Order
 
	err := db.Preload("Items").First(&order, id).Error
 
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": fmt.Sprintf("user with id %v not found", id),
		})
		return
	}
 
	if err = c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		return
	}
 
	for _, item := range order.Items {
		updatedItem := models.Item {
			Item_code: item.Item_code,
			Description: item.Description,
			Quantity: item.Quantity,
		}
 
		err = db.Model(&item).Where("order_id = ?", item.Item_id).Updates(updatedItem).Error
 
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
			return
		}
	}
 
	err = db.Model(&order).Where("order_id = ?", id).Updates(order).Error
 
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		return
	}
 
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": order,
	})
}

// DeleteData
func DeleteData(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
 
	var order models.Order
 
	err := db.Preload("Items").First(&order, id).Error
 
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": fmt.Sprintf("user with id %v not found", id),
		})
		return
	}
 
	for _, item := range order.Items {
		db.Delete(&item)
	}
 
	db.Delete(&order)
 
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("order with id %v has been successfully deleted", id),
	})
}
