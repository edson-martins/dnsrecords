[![Go](https://github.com/edson-martins/dnsrecords/actions/workflows/go.yml/badge.svg?style=plastic)](https://github.com/edson-martins/dnsrecords/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/edson-martins/dnsrecords?style=plastic)](https://goreportcard.com/report/github.com/edson-martins/dnsrecords)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=edson-martins_dnsrecords&metric=alert_status)](https://sonarcloud.io/dashboard?id=edson-martins_dnsrecords)
---

<p align="center">
    <img src="https://img.shields.io/badge/Go-1.16.3-informational?style=plastic&logo=go&logoColor=white&color=7FD5E9"/> <img src="https://img.shields.io/badge/Mux-1.8.0-informational?style=plastic&logo=mux&logoColor=white&color=7FD5E9"/> <img src="https://img.shields.io/badge/go--logging-v1-informational?style=plastic&logo=mux&logoColor=white&color=7FD5E9"/>
</p>

# DNS MX Record Web Application

A web application that displays the IP addresses and mail hosts associated with a user-entered domain name.

## Requirements
· The application must listen on port 8080

· When accessing the root app (http://localhost:8080/), display a form asking for a domain name (e.g. domain.com)

· When that value is submitted, display the IP addresses associated with the domain and the hosts associated with that domain's DNS MX records.

· If the results page is accessed with the HTTP Accept header set to "application/json", render a JSON response instead of HTML

    o Note, in this case, the Accept header will only be "application/json". 

## Installation

````
$ go get github.com/edson-martins/dns-mx-record
$ cd dns-mx-record
$ go mod init
````
### Prerequisites

 This application is using the libraries below as dependency:

 1 - Golang logging library - github.com/op/go-logging
 ```
 $ go get github.com/op/go-logging
 ```
 2 - gorilla/mux - github.com/gorilla/mux
 ```
 $ go get -u github.com/gorilla/mux
 ```

## Getting Started

 1 - There is a contants.go file used to store the URL Port (default is 8080) and the Index template path. If you want to change the host port, that's the place.

 
 2 - The program can be executed in the project root path (path where the main.go file is available) using the command below:
 ```
 go run main.go
 ```
 3 - The build process looking to generate the executable can be done as:
 ```
 go build
 go install
 ```
 The main file executable will be installed on $GOPATH/bin and the dependecies at $GOPATH/pkg