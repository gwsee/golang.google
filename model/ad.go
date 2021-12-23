package model

type Ad struct {
	Id           string   `json:"id"`
	ResourceName string   `json:"resourceName"`
	FinalUrls    []string `json:"finalUrls"`
}
