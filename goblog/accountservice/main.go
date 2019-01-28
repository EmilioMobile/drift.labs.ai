package main

import (
	"fmt"
	service "github.com/EmilioMobile/drift.labs.ai/goblog/accountservice/services"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}
