package controllers

import (
  "github.com/gin-gonic/gin"
  "os/exec"
  "net/http"
  "io/ioutil"
  "strings"
  "log"
  "os"
  "bytes"
  "fmt"
)

type GenerateMaintenancePlanInput struct {
  MaintenanceType  string `json:"maintenanceType" binding:"required"`
  Name string `json:"name" binding:"required"`
  ApplicationCode string `json:"applicationCode" binding:"required"`
  Recurrence string `json:"recurrence" binding:"required"`
  RecurrenceDayOfWeek string `json:"recurrence" binding:"required"`
  RecurrenceDayOfMonth string `json:"recurrence" binding:"required"`
  StartDate string `json:"startDate" binding:"required"`
  EndDate string `json:"endDate" binding:"required"`
}

// POST /generate_maintenance_plan

func GenerateMaintenancePlan(c *gin.Context) {

    var input GenerateMaintenancePlanInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    logfile := "generateMaintenancePlan.log"
    maintenance_plan_csv := "/root/dynatrace-script-configuration/Create_new_application/Dynatrace_Maintenance_Plan.csv"

    f, err := os.OpenFile(logfile,
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()

    logger := log.New(f, "prefix", log.LstdFlags)

    logger.Println("############################################################################################")
    logger.Println("API Call for Maintenance Window creation : ")
    logger.Println("############################################################################################")

    values := input.ApplicationCode+";"+input.MaintenanceType+";"+input.Name+";"+input.Recurrence+";"+input.RecurrenceDayOfWeek+";"+input.RecurrenceDayOfMonth+";"+input.StartDate+";"+input.EndDate+";"
    logger.Println("API Call parameters "+values)

    // read the whole file at once
    b, err := ioutil.ReadFile(maintenance_plan_csv)
    if err != nil {
        logger.Println("An occured while opening the file ")
        logger.Println(err.Error())
        return
    }
    
    // check whether s contains substring text
    s := string(b)

    isPresent := strings.Contains(s, values)

    if isPresent {
        logger.Println("API Call for an already existing maintenance plan "+values)
        c.JSON(http.StatusOK, gin.H{"data": "Maintenance Plan already exists"})
        return
    }

    file, err := os.OpenFile(maintenance_plan_csv, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        logger.Println("An occured while opening the file "+maintenance_plan_csv)
        logger.Println(err.Error())
    }

    defer file.Close()
    if _, err := file.WriteString(values); err != nil {
        logger.Println("An occured while opening the file "+maintenance_plan_csv)
        logger.Println(err)
    }

    cmd := exec.Command("python3","/root/dynatrace-script-configuration/Create_new_application/4_Generate_Maintenance_Plan.py")
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
    
    logger.Println("Maintenance window successfully created "+values)

    c.JSON(http.StatusOK, gin.H{"data": "Maintenance window successfully created"})
}
