package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteDataToFileSuccess(t *testing.T) {
	var mockData = []byte(`{
		"FirstName": "Guru",
		"LastName": "Singh"
	   }`)
	err := WriteDataToFile(mockData)
	assert.Nil(t, err)
}

func TestWriteDataToFileFail(t *testing.T) {
	var mockData = []byte{}
	err := WriteDataToFile(mockData)
	assert.NotNil(t, err)
}
