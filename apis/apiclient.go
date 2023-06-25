package apis

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tech-mahindra/public-apis/config"
	"github.com/tech-mahindra/public-apis/file"
	"github.com/tech-mahindra/public-apis/utils"
)

type apiClient struct {
	endPoints []string
	batchSize int // batchSize for how many api will hit concurrently
	urlBatch  chan string
	dataChan  chan []byte
	totalData []byte
	Duration  time.Duration // Duration in hours how often data is required
}

var client *apiClient
var clientConfig = config.GetClientConfig()

func init() {
	client = &apiClient{
		endPoints: clientConfig.EndPoints,
		batchSize: clientConfig.Concurency,
		dataChan:  make(chan []byte, clientConfig.Concurency),
		Duration:  1,
		urlBatch:  make(chan string, clientConfig.Concurency),
	}
}

// GetApiClient -
func GetApiClient() *apiClient {
	return client
}

// GetApiClient - method to start scheduler
func (c *apiClient) StartJob() {
	go c.recieveDataFromChannel()
	go c.sendDataToChannel()
}

// sendDataToChannel - method to send the data to channel recieved from the apis
func (c *apiClient) sendDataToChannel() {
	for len(c.endPoints) != 0 {
		url := c.endPoints[0]
		c.urlBatch <- url
		c.endPoints = c.endPoints[1:]
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered Error in fetchData:\n", r)
					c.dataChan <- nil
					return
				}
			}()
			data, err := fetchData(url)
			if err != nil {
				panic(err)
			}
			c.dataChan <- data
		}()
	}

}

// recieveDataFromChannel - method to recieve the data from channel
func (c *apiClient) recieveDataFromChannel() {
	var responseCount = 0
	for {
		if len(c.totalData) > 0 {
			c.totalData = utils.MergeSlices(c.totalData, <-c.dataChan)
		} else {
			c.totalData = append(c.totalData, <-c.dataChan...)
		}
		responseCount++
		fmt.Println("urls processed ", <-c.urlBatch)
		if responseCount == len(clientConfig.EndPoints) {
			fmt.Println("all api response recieved = ", len(c.totalData))
			err := file.WriteDataToFile(c.totalData)
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(c.Duration * time.Hour)
			fmt.Println("Restarting schduler = ", len(c.totalData))
			c.endPoints = clientConfig.EndPoints
			c.totalData = []byte{}
			responseCount = 0
			go c.sendDataToChannel()
		}
	}
}

// fetchData - method to get the data from api
func fetchData(url string) (data []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	return data, err
}
