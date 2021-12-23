package model

type Campaign struct {
	Id                     string `json:"id"`
	Name                   string `json:"name"`
	ResourceName           string `json:"resourceName"`
	AdvertisingChannelType string `json:"advertisingChannelType"`
}
