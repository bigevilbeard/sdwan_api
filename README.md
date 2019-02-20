# SDWAN API Example w/ Golang

## Introduction

The Cisco SDWAN software provides a REST API, which is a programmatic interface for controlling, configuring, and monitoring the Cisco SDWAN devices in an overlay network. You access the REST API through the vManage web server.

A REST API is a web service API that adheres to the REST, or Representational State Transfer, architecture. The REST architecture uses a stateless, client–server, cacheable communications protocol. The Cisco SDWAN vManage NMS web server uses HTTP and its secure counterpart, HTTPS, as the communications protocol. REST applications communicate over HTTP or HTTPS using standard HTTP methods to make calls between network devices. These standard HTTP methods include:

- GET—Retrieve or read information.
- PUT—Update an object.
- POST—Create an object.
- DELETE—Remove an object.

REST is a simpler alternative to mechanisms such as remote procedure calls (RPCs) and web services such as Simple Object Access Protocol (SOAP) and Web Service Definition Language (WSDL). The REST architecture has not been formally defined by any standards bodies.

For more information about REST, see this [Wikipedia](https://en.wikipedia.org/wiki/Representational_state_transfer) page.

## Resource Data
The Cisco SDWAN REST API uses the JavaScript Object Notation (JSON) data model to represent the data associated with a resource. JSON has three types of data:

- Scalar—Data types with a single value. Scalar types include numbers, strings, and booleans.
- Array—Ordered lists of values of an arbitrary type.
- Object—Unordered set of key:value pairs (also referred to as attributes), where key is a string and the value can have an arbitrary type.

For more information about JSON, see the JSON Data [Interchange Standard](http://json.org/).


```
git clone https://github.com/bigevilbeard/sdwan_api
```

## Install go(lang)
With [homebrew](http://mxcl.github.io/homebrew/):
```
sudo brew install go
```
With [apt](http://packages.qa.debian.org/a/apt.html)-get:
```
sudo apt-get install golang
```
[Install Golang](https://golang.org/doc/install) manually or [compile](https://golang.org/doc/install/source) it yourself

Please follow the instructions [here](https://golang.org/doc/install) if you need to install Go and are unsure, would like to see the documents

## 'GET' this code

All of the code and examples for this example can be cloned and access it with the following commands:

```
git clone https://github.com/bigevilbeard/sdwan_api
```

The code is build to be used on the Cisco DevNet always on [SDWAN Sandbox](https://devnetsandbox.cisco.com/RM/Diagram/Index/4fb544ad-c88c-4227-8b09-5d35aa26a63b?diagramType=Topology) - to make this work in another environment update the `baseURL` / `Username/Password` in the code

## The Code

When i built this code, i broke it down into sections and steps, below is how i did this. With each example being a checkpoint and the code following this the next step. This helped me build this code.

- `example.api.cookie` - When you use a program or script to transfer data from a vManage web server or perform operations on the server, you must first establish an HTTPS session to the server. To do this, you send a call to log in to the server with the following parameters:

- URL to send the request to—Use `https://{vmanage-ip-address/j_security_check`, which performs the login operation and security check on the vManage web server at the specified IP address.
- Request method—Specify a `Post` request.
- The data contains the username and password in the format `j_username=username` & `j_password=password`.

Note: Disable security check for HTTPS(SSL) for certificate have been implemented see:[import "crypto/tls"](https://golang.org/pkg/crypto/tls/) as this is sandbox.

I added the HTTP Response Status Codes. When making http requests with Go it is almost always necessary to check the status code of the response which is returned. Generally, if the status code is between 200 and 300 you can treat as successful. But anything except a 200-300 status, we often need to fix something.

- `example.api.func` This breaks the `example.api.cookie` into two functions `main()` function will now utilize all the functions creating a `cookiejar` to store cookies required while logging into the URL/vManage

- `example.api.output` adds an additional `Get` request for `https://{vmanage-ip-address/dataservice/device` this display all Cisco SDWAN devices in the overlay network that are connected to the vManage NMS. The output is

- `example.api.final`Printing the final output into a more human readable format and removing unwanted information the [import "encoding/json"](https://golang.org/pkg/encoding/json/) has been used. [interface](https://golang.org/pkg/fmt/) make the code more flexible, scalable and it’s a way to achieve [polymorphism](https://en.wikipedia.org/wiki/Polymorphism_%28computer_science%29) in Go. Instead of requiring a particular type, interfaces allow to specify that only some behaviour is needed.


## About me

Network Automation Developer Advocate for Cisco Devnet.
I'm like Hugh Hefner... minus the mansion, the exotic cars, the girls, the magazine and the money. So basically, I have a robe.

Find me here: [LinkedIn](https://www.linkedin.com/in/stuarteclark/) / [Twitter](https://twitter.com/bigevilbeard)
