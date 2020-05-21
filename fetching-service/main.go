package main

import (
	"fmt"
	"log"
	"os"

	"github.com/efishery/fetchin-service/cache"
	"github.com/efishery/fetchin-service/docs"
	routes "github.com/efishery/fetchin-service/module/fetch/routes"
)

var (
	err error
)

func main() {
	command := parseCommand()
	switch command {
	case "serve":
		startApp()
	default:
		fmt.Printf("Invalid command %s\n", command)
	}

}

func startApp() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Efishery API - Fetching Service"
	docs.SwaggerInfo.Description = "This is an API for Efishery Fetching service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = ""
	cache, err := cache.NewCacheDefault()
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	r := routes.SetupRouter(cache)
	//running server
	r.Run(":9001")
}

func parseCommand() string {
	args := os.Args[1:]
	if len(args) > 0 {
		return args[0]
	}
	return ""
}
