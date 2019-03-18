package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

type List struct {
    Status string   `json:"status"`
    Data   []string `json:"data"`
}

type Query struct {
    Status string `json:"status"`
    Data   struct {
        ResultType string `json:"resultType"`
        Result     []struct {
            Metric struct {
                Name     string `json:"__name__"`
                Instance string `json:"instance"`
                Job      string `json:"job"`
                Version  string `json:"version"`
            } `json:"metric"`
            Value []interface{} `json:"value"`
        } `json:"result"`
    } `json:"data"`
}

func main() {
    flag.Parse()
    args := flag.Args()
    if len(args) == 0 {
        os.Exit(1)
    }
    fmt.Println("Host:", args[0])

    localhost := "127.0.0.1:9090"
    url := "http://"
    if len(args) == 1 && args[0] != "ls" {
        url = url + localhost + "/api/v1/query?query=" + args[0]
    } else if len(args) == 2 && args[1] != "ls" {
        url = url + args[0] + "/api/v1/query?query=" + args[1]
    } else if args[0] == "ls" {
        url = url + localhost + "/api/v1/label/__name__/values"
    } else if args[1] == "ls" {
        url = url + args[0] + "/api/v1/label/__name__/values"
    }

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    barray, _ := ioutil.ReadAll(resp.Body)
    jsonBytes := []byte(barray)
    data := new(List)
    if err = json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(data.Data[:])
}
