package main

import (
  "github.com/gin-gonic/gin"
  "controllers"
)

func main() {
  r := gin.Default()

  r.GET("/", controllers.Welcome)
  r.POST("/generateConfiguration", controllers.GenerateConfiguration)
  r.GET("/updateHomeDashboard", controllers.UpdateHomeDashboard)
  r.POST("/pgAvailability", controllers.PGAvailability)
  r.POST("/generateMaintenancePlan", controllers.GenerateMaintenancePlan)
  
  r.Run("l203pdo017:7777")
}
