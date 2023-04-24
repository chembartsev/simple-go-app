package main

import (
	"fmt"
	"net/http"
	"simple-go-app/config"
	"simple-go-app/metrics"
	"time"
)

func main() {
	configData := config.GetConfig()

	// Starting the server
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", configData.Metrics.Port), nil)
		if err != nil {
			panic(err)
		}
	}()

	// Just to see if the app is alive
	go func() {
		for {
			time.Sleep(time.Duration(configData.StdOutLogTimeout) * time.Millisecond)
			fmt.Printf("[%s] %s is alive\n", time.Now().UTC().Format(time.RFC3339), configData.ProjectName)
		}
	}()

	// Starting message
	fmt.Printf("[%s] %s is started\n", time.Now().UTC().Format(time.RFC3339), configData.ProjectName)

	metrics.RegisterMetrics(metrics.NewMetrics())
	select {}
}
