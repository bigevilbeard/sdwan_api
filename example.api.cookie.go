package main

import (
  "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"crypto/tls"
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
type AuthenticityToken struct {
	Token string
}

func main() {
	jar, _ := cookiejar.New(nil)
	tr := &http.Transport{
			 TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	app := App{
		Client: &http.Client{Jar: jar,Transport: tr},
	}
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

	bodyBytes, err := ioutil.ReadAll(response.Body)
  bodyString := string(bodyBytes)
  fmt.Println(bodyString)
	fmt.Println("HTTP Response Status:", response.StatusCode, http.StatusText(response.StatusCode))

  if response.StatusCode >= 200 && response.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")
  } else {
		fmt.Println("Argh! Broken")
 }
}
