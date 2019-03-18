package main

import (
  "fmt"
  "flag"
  "os"
  "net/http"
  "io/ioutil"
)

func main() {
  flag.Parse()
  args := flag.Args()
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
  fmt.Println(string(barray))
}
