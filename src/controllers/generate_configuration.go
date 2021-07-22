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

type GenerateConfigurationInput struct {
  Tenant  string `json:"tenant" binding:"required"`
  Env string `json:"env" binding:"required"`
  ApplicationCode string `json:"applicationCode" binding:"required"`
  Name string `json:"name" binding:"required"`
  MZType string `json:"mztype" binding:"required"`
  Domain string `json:"domain" binding:"required"`
  Synthetic string `json:"synthetic" binding:"required"`
  GOOasis string `json:"gooasis" binding:"required"`
  OracleInsight string `json:"oracleinsight" binding:"required"`
  Kubernetes string `json:"kubernetes" binding:"required"`
}

// POST /generate_configuration

func GenerateConfiguration(c *gin.Context) {

    var input GenerateConfigurationInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    //if pas bon tenant return
    //if pas bon env return
    //if pas bon code appli return
    //if mztype pas bon return
    //if gooasis pas date format return
    //if oracleinsight pas booleen
    //if kubernetes pas booleen

    logfile := "GenerateConfiguration.log"
    configuration_csv := "/root/dynatrace-script-configuration/Create_new_application/Dynatrace_Appli_Stime_v2.csv"

    f, err := os.OpenFile(logfile,
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()

    logger := log.New(f, "prefix", log.LstdFlags)

    logger.Println("############################################################################################")
    logger.Println("API Call for Configuration creation : ")
    logger.Println("############################################################################################")

    values := input.Tenant+";"+input.Env+";"+input.ApplicationCode+";"+input.Name+";"+input.MZType+";"+input.Domain+";"+input.Synthetic+input.GOOasis+";"+input.OracleInsight+";"+input.Kubernetes
    logger.Println("API Call parameters "+values)

    // read the whole file at once
    b, err := ioutil.ReadFile(configuration_csv)
    if err != nil {
        logger.Println("An occured while opening the file ")
        logger.Println(err.Error())
        return
    }
    
    // check whether s contains substring text
    s := string(b)

    isPresent := strings.Contains(s, values)

    if isPresent {
        logger.Println("API Call for an already existing configuration "+values)
        c.JSON(http.StatusOK, gin.H{"data": "Configuration already exists"})
        return
    }

    file, err := os.OpenFile(configuration_csv, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        logger.Println("An occured while opening the file "+configuration_csv)
        logger.Println(err.Error())
    }

    defer file.Close()
    if _, err := file.WriteString(values); err != nil {
        logger.Println("An occured while opening the file "+configuration_csv)
        logger.Println(err)
    }

    cmd := exec.Command("python3","/root/dynatrace-script-configuration/Create_new_application/1_Generate_Config_STIME_v2.1.py")
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
    
    logger.Println("Configuration successfully created "+values)

    c.JSON(http.StatusOK, gin.H{"data": "Configuration successfully created "})
}
