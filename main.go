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
    TYPE string
    CMD string
    AMT int64
}  
    



func main() {

    Ecr := Ecr{TYPE: "EDC",CMD:"SALE",AMT: 1,}
    
    jsonstring, _ := json.Marshal(Ecr)
    fmt.Print(string(jsonstring)+"\n")

    resp,_ := http.Post("http://172.16.104.91:8080/", "application/json", bytes.NewBuffer(jsonstring))

    defer resp.Body.Close()

    body, _:=ioutil.ReadAll(resp.Body)

    fmt.Println(string(body))
    }