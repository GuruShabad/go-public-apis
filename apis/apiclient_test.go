package apis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const mockUrl = "https://cat-fact.herokuapp.com/facts"

var mockData = []byte{20, 30}

func Test_GetApiClientSuccess(t *testing.T) {
	got := GetApiClient()
	assert.NotNil(t, got, "GetApiClient")
	assert.Equal(t, client, got)
}

func Test_GetApiClientFail(t *testing.T) {
	client := &apiClient{}
	got := GetApiClient()
	assert.NotNil(t, got, "GetApiClient")
	assert.NotEqual(t, client, got)
}

func Test_sendDataToChannelSuccess(t *testing.T) {
	client.endPoints = []string{mockUrl}
	client.sendDataToChannel()
	res := <-client.dataChan
	assert.NotNil(t, res, "Test_sendDataToChannel")
}

func Test_sendDataToChannelFail(t *testing.T) {
	client.endPoints = []string{"url"}
	client.sendDataToChannel()
	res := <-client.dataChan
	assert.Nil(t, res, "Test_sendDataToChannel")
}

func Test_recieveDataFromChannelSuccess(t *testing.T) {
	client.dataChan <- mockData
	go client.recieveDataFromChannel()
	res := <-client.dataChan
	assert.NotNil(t, res, "Test_recieveDataFromChannelSuccess")

}

func Test_fetchDataSuccess(t *testing.T) {
	resp, err := fetchData(mockUrl)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

}

func Test_fetchDataFail(t *testing.T) {
	mockUrl := "https://"
	resp, err := fetchData(mockUrl)
	assert.NotNil(t, err)
	assert.Nil(t, resp)

}
