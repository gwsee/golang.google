package model

type Metrics struct {
	CostMicros                     int     `json:"costMicros,string"`
	ActiveViewMeasurableCostMicros int     `json:"activeViewMeasurableCostMicros,string"`
	Clicks                         int     `json:"clicks,string"`
	InvalidClicks                  int     `json:"invalidClicks,string"`
	InvalidClickRate               float64 `json:"invalidClickRate"`
	Impressions                    int     `json:"impressions,string"`
	ImpressionsRate                float64 `json:"interaction_rate"`
}
