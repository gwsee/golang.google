package service

// GetAdGroupAdList
var sql001 = `select campaign.id,campaign.name,campaign.status,campaign.advertising_channel_type,
				metrics.impressions,metrics.cost_micros,
				ad_group.id,ad_group_ad.ad.id,ad_group_ad.ad.final_urls,segments.date 
              from ad_group_ad 
				where segments.date BETWEEN '%v' AND '%v' AND metrics.cost_micros>0 ORDER BY segments.date`
