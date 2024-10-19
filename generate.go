package ksperimeterxsdkgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (r *PxRequest) RequestPx() (*PxResponse, error) {
	requestData, err := structToReader(r)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", ApiUrl, requestData)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var responseData PxResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}
	if !responseData.Success {
		return nil, errors.New("could not successfully generate px")
	}
	return &responseData, err
}

func structToReader(data interface{}) (io.Reader, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(jsonData), nil
}
