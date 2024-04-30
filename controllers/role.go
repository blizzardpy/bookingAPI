package controllers

import (
	"bookingAPI/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRole(c *gin.Context) {
	var Role models.Role
	c.BindJSON(&Role)
	err := models.CreateRole(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Role)
}

func GetRoles(c *gin.Context) {
	var Roles []models.Role
	err := models.GetRoles(&Roles)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Roles)
}

func GetRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var Role models.Role
	err := models.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

func UpdateRole(c *gin.Context) {
	var Role models.Role
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&Role)
	err = models.UpdateRole(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Role)
}
