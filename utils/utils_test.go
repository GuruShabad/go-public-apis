package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSlices(t *testing.T) {
	var mockData1 = []byte(`[{
		"FirstName": "Guru",
		"LastName": "Singh"
	   }]`)

	var mockData2 = []byte(`[{
		"FirstName": "Guru",
		"LastName": "Singh"
	   }]`)

	mergedData := MergeSlices(mockData1, mockData2)
	assert.NotNil(t, mergedData, "TestMergeSlices")

}
