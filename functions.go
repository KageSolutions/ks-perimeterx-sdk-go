package ksperimeterxsdkgo

import (
	"bytes"
	"encoding/json"
	"io"
)

// init new sdk instance
func NewPxSdkInstance(apiKey string, site PxSite) *PxSdkInstance {
	return &PxSdkInstance{
		request: pxRequest{ApiKey: apiKey, Site: site},
	}
}

func (p *PxSdkInstance) UpdatePxHdCookie(pxHd string) {
	p.request.PXHD = pxHd
}

func (p *PxSdkInstance) UpdateProxy(proxy string) {
	p.request.Proxy = proxy
}

func (p *PxSdkInstance) UpdatePageUrl(pageURL string) {
	p.request.PageURL = pageURL
}

func (p *PxSdkInstance) UpdateUserAgent(userAgent string) {
	p.request.Ua = userAgent
}

func structToReader(data interface{}) (io.Reader, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(jsonData), nil
}
