package controllers

import (
  "github.com/gin-gonic/gin"
  "os/exec"
  "net/http"
  "log"
  "os"
  "bytes"
  "fmt"
)

type PGAvailabilityInput struct {
  Env string `json:"env" binding:"required"`
  ApplicationCode string `json:"applicationCode" binding:"required"`
  Name string `json:"name" binding:"required"`
  Domain string `json:"domain" binding:"required"`
}

// POST /pg_availability

func PGAvailability(c *gin.Context) {

    var input PGAvailabilityInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    //if pas bon env return
    //if pas bon code appli return

    logfile := "PGAvailability.log"
    
    f, err := os.OpenFile(logfile,
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()

    logger := log.New(f, "prefix", log.LstdFlags)

    logger.Println("############################################################################################")
    logger.Println("API Call for PG Availability : ")
    logger.Println("############################################################################################")

    values := input.Env+";"+input.ApplicationCode+";"+input.Name+";"+input.Domain
    logger.Println("API Call parameters "+values)

    cmd := exec.Command("python3","/root/dynatrace-script-configuration/Create_new_application/3_PGAvailability.py")
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Dir = "/root/dynatrace-script-configuration/Create_new_application/"
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err_script := cmd.Run()
    if err_script != nil {
        logger.Println("Error occured during script execution")
        logger.Println(stderr.String())
        return
    }
    
    logger.Println("PG Availability successfully executed "+values)

    c.JSON(http.StatusOK, gin.H{"data": "PG Availability successfully executed"})
}
