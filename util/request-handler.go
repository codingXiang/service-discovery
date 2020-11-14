package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func PostJSON(url string, body interface{}) error {
	jsonStr, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 {
		return nil
	} else {
		return errors.New("status code is not equal 200, is " + strconv.Itoa(resp.StatusCode))
	}
}
