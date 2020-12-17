
package main

import (
	"github.com/k8s-go-demo/app"
	"github.com/k8s-go-demo/config"
	"log"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	log.Println("Listening on https://localhost:8800/")
	app.Run(":8800")
}
