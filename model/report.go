package model

type AdGroupAdList = []AdGroupAdData

type AdGroupAdData struct {
	Results   []AdGroupAdInfo `json:"results"`
	FieldMask string          `json:"fieldMask"`
	RequestId string          `json:"requestId"`
}

type AdGroupAdInfo struct {
	Campaign  Campaign  `json:"campaign"`
	AdGroup   AdGroup   `json:"adGroup"`
	AdGroupAd AdGroupAd `json:"adGroupAd"`
	Segments  Segments  `json:"segments"`
	Metrics   Metrics   `json:"metrics"`
}
