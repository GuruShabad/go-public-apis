package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// WriteDataToFile - function to write the data to the file
func WriteDataToFile(data []byte) (err error) {
	now := time.Now() // current local time
	sec := now.Unix()
	fileName := fmt.Sprintf("puclic-api_%d.json", +sec)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, data, "", "  "); err != nil {
		return err
	}

	_, err = file.Write(dst.Bytes())
	return err
}
