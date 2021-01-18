package main

import (
    "fmt"
    "encoding/json"
    // "os"
    // "crypto/md5"
    "net/http"
    "io/ioutil"
    // "strings"
    "bytes"
    // "log"
)
type Ecr struct {
    TYPE string  `json:"TYPE"`
    CMD string `json:"CMD"`
    AMT float64 `json:"AMT"`
    }

func main() {
    
     var TYPENAME string
     var CMDNAME string
     var ECRAMT float64

    fmt.Println("Enter ECR TYPE:")
    fmt.Scanln(&TYPENAME)

    fmt.Println("Enter ECR CMD:")
    fmt.Scanln(&CMDNAME)

    fmt.Println("Enter ECR AMT:")
    fmt.Scanln(&ECRAMT)

    Ecr := &Ecr{TYPE:TYPENAME,CMD:CMDNAME,AMT:ECRAMT}

    jsonstring, _ := json.Marshal(Ecr)
    fmt.Print(string(jsonstring)+"\n")

    resp,_ := http.Post("http://172.16.104.27:8080/", "application/json", bytes.NewBuffer(jsonstring))

    defer resp.Body.Close()

    body, _:=ioutil.ReadAll(resp.Body)

    fmt.Println(string(body))
    }