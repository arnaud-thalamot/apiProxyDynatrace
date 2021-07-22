package controllers

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

// GET /Welcome

func Welcome(c *gin.Context) {

    c.JSON(http.StatusOK, gin.H{"data": "Welcome to DynAPI"})
}