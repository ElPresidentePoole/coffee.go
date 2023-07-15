package main

import (
  "log"
  "fmt"
  "net/http"
  "os/exec"
)

func GetPageHTML(coffeeMakerOn bool) string {
  var coffeeStatus string
  if coffeeMakerOn {
    coffeeStatus = "On ☕"
  } else {
    coffeeStatus = "Off 😴"
  }
  return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>coffee.go</title>
  </head>
  <body>
  <p>Status: %s!</p>
  <p><a href="./coffeeoff">Off</a></p>
  <p><a href="./coffeeon">On</a></p>
  </body>
</html>`, coffeeStatus)
}

func main() {
  coffeeStatus := false
  pin := 7
  http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(rw, "%s", GetPageHTML(coffeeStatus))
  })
  http.HandleFunc("/coffeeon", func(rw http.ResponseWriter, r *http.Request) {
    coffeeStatus = true
    fmt.Fprintf(rw, "%s", GetPageHTML(coffeeStatus))
    log.Println("Coffee maker turned on!")
    cmd := exec.Command("gpio", "mode", fmt.Sprint(pin), "up")
    _, err := cmd.Output()
    if err != nil {
      fmt.Println(err.Error())
    }
  })
  http.HandleFunc("/coffeeoff", func(rw http.ResponseWriter, r *http.Request) {
    coffeeStatus = false
    fmt.Fprintf(rw, "%s", GetPageHTML(coffeeStatus))
    log.Println("Coffee maker turned off!")
    cmd := exec.Command("gpio", "mode", fmt.Sprint(pin), "down")
    _, err := cmd.Output()
    if err != nil {
      fmt.Println(err.Error())
    }
  })

  log.Fatal(http.ListenAndServe(":1337", nil))
}
