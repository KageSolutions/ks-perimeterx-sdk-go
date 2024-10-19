package ksperimeterxsdkgo

var (
	ApiUrl = "https://perimeterx.prod.kageapis.com/web/cookie/init"
)

const (
	WalmartUS  = PxSite("walmartus")
	WalmartCA  = PxSite("walmartca")
	Hibbett    = PxSite("hibbett")
	VividSeats = PxSite("vividseats")
	StockX     = PxSite("stockx")
)

type PxSite string
