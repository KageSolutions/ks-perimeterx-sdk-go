package ksperimeterxsdkgo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// request px init cookie
func (r *PxSdkInstance) RequestPxInit() (*PxResponse, error) {
	requestData, err := structToReader(r.request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", ApiInitUrl, requestData)
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
	r.request.Data = responseData.Headers.Data
	return &responseData, err
}
