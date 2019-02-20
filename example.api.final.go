package main

import (
  "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"crypto/tls"
  "encoding/json"
)
const (
	baseURL = "https://sandboxsdwan.cisco.com:8443"
)

var (
	username = "devnetuser"
	password = "Cisco123!"
)

type App struct {
	Client *http.Client
}

func (app *App) login()  {
  client := app.Client
  loginURL := baseURL + "/j_security_check"
  data := url.Values{
      "j_username": {username},
      "j_password": {password},
  }
  response, err := client.PostForm(loginURL, data)
  if err != nil {
     log.Fatal(err)
  }
  defer response.Body.Close()

  fmt.Println("HTTP Response Status:", response.StatusCode, http.StatusText(response.StatusCode))
  if response.StatusCode >= 200 && response.StatusCode <= 299 {
    fmt.Println("HTTP Status is in the 2xx range")
    } else {
      fmt.Println("Argh! Broken")
}
}
func (app *App) getDevices() {
        client := app.Client
        devicesURL := baseURL + "/dataservice/device"
        response, err := client.Get(devicesURL)

        if err != nil {
           log.Fatal(err)
        }
        defer response.Body.Close()

        if err != nil {
           log.Fatal(err)
        }
        body, _ := ioutil.ReadAll(response.Body)


        var data1 map[string]interface{}
	      jsonErr := json.Unmarshal(body, &data1)
	      if jsonErr != nil {
		             log.Fatal(jsonErr)
        }


        devices := data1["data"].([]interface{})

        fmt.Println("------")
        for value := range devices {
          // Assigning unknown
          current_device := devices[value].(map[string]interface{})

          fmt.Println("| Hostname |", current_device["host-name"],"| Device ID |", current_device["deviceId"],
          "| Reachability |", current_device["reachability"],"| Version |", "| Version |", current_device["version"])

        }


}
func main() {
        jar, _ := cookiejar.New(nil)
        tr := &http.Transport{
          TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	      }
	      app := App{
          Client: &http.Client{Jar: jar,Transport: tr},
        }
        app.login()
        app.getDevices()
 }
