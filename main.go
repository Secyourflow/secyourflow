package main

import (
	"log"

	"github.com/lagzenthakuri/secyourflow/app"
)

func main() {
	secyourflow, err := app.App("./secyourflow_data", false)
	if err != nil {
		log.Fatal(err)
	}

	if err := secyourflow.Start(); err != nil {
		log.Fatal(err)
	}
}
