package controllers

import (
	"github.com/gin-gonic/gin"
)

func HelloWorlder(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World from k8s!"})
}
