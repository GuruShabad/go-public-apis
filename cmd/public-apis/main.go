package main

import (
	"sync"

	"github.com/tech-mahindra/public-apis/apis"
	"github.com/tech-mahindra/public-apis/config"
)

var logger *config.Logger

func init() {
	logger = config.GetLogger()
}

func main() {
	logger.Info.Println("Starting the application...")

	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()

	apiClient := apis.GetApiClient()
	go apiClient.StartJob()

}
