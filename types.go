package ksperimeterxsdkgo

type PxRequest struct {
	ApiKey  string `json:"apiKey"`
	Ua      string `json:"ua"`
	PageURL string `json:"pageUrl"`
	Proxy   string `json:"proxy"`
	Cookie  string `json:"cookie"`
	PXHD    string `json:"_pxhd"`
	Data    string `json:"data,omitempty"`
}

type PxResponse struct {
	Success bool   `json:"success"`
	Cookie  string `json:"cookie"`
	Cts     string `json:"cts"`
	Vid     string `json:"vid"`
	Headers struct {
		SecUa     string `json:"secUa"`
		Platform  string `json:"platform"`
		UserAgent string `json:"userAgent"`
	} `json:"headers"`
}
