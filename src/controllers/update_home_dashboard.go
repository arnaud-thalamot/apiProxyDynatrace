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

// GET /update_home_dashboard

func UpdateHomeDashboard(c *gin.Context) {

    logfile := "UpdateHomeDashboard.log"

    f, err := os.OpenFile(logfile,
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()

    logger := log.New(f, "prefix", log.LstdFlags)

    logger.Println("############################################################################################")
    logger.Println("API Call for UpdateHomeDashboard : ")
    logger.Println("############################################################################################")

    cmd := exec.Command("python3","/root/dynatrace-script-configuration/Create_new_application/2_Update_Home_Dashboard.py")
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
    
    logger.Println("Home Dashboard successfully updated ")

    c.JSON(http.StatusOK, gin.H{"data": "Home Dashboard successfully updated "})
}
