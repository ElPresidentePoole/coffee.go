package main

import (
  "log"
  "fmt"
  "net/http"
  "github.com/stianeikeland/go-rpio/v4"
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

var coffeeMakerPin = rpio.Pin(13)

func turnOnCoffeeMaker() {
  coffeeMakerPin.High()
}

func turnOffCoffeeMaker() {
  coffeeMakerPin.Low()
}

func main() {
  if err := rpio.Open(); err != nil {
    log.Fatal(err)
  }
  defer rpio.Close()
  coffeeMakerPin.Output()

  coffeeStatus := false
  http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(rw, "%s", GetPageHTML(coffeeStatus))
  })
  http.HandleFunc("/coffeeon", func(rw http.ResponseWriter, r *http.Request) {
    coffeeStatus = true
    fmt.Fprintf(rw, "%s", GetPageHTML(coffeeStatus))
    log.Println("Coffee maker turned on!")
    turnOnCoffeeMaker()
  })
  http.HandleFunc("/coffeeoff", func(rw http.ResponseWriter, r *http.Request) {
    coffeeStatus = false
    fmt.Fprintf(rw, "%s", GetPageHTML(coffeeStatus))
    log.Println("Coffee maker turned off!")
    turnOffCoffeeMaker()
  })

  log.Fatal(http.ListenAndServe(":1337", nil))
}
