package main

import (
  "fmt"
  "flag"
  "net/http"
  "io/ioutil"
)

func main() {
  flag.Parse()
  args := flag.Args()
  fmt.Println("Host:", args[0])

  url := "http://" + args[0] + "/api/v1/query?query=go_info"
  resp, _ := http.Get(url)
  defer resp.Body.Close()

  barray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(barray))
}
