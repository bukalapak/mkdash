package main

import (
	"log"
	"os"

	"github.com/bukalapak/mkdash/maker"
	"github.com/bukalapak/mkdash/parser"
)

var (
	GrafanaAuth    string
	grafanaURL     = "https://grafana-mon.bldevs.info"
	grafanaTimeOut = 2
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("format should be `mkdash [project-name] [service-name]")
	}
	data := &parser.InputData{
		Project: args[0],
		Service: args[1],
	}

	pr := parser.New()
	mk := maker.New(grafanaURL, GrafanaAuth, grafanaTimeOut)

	body, err := pr.ParseTemplate(data)
	if err != nil {
		log.Fatal(err)
	}

	err = mk.CreateDashboard(body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfuly create dashboard")
}
