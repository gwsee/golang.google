package http

import "git.zx-tech.net/pengfeng/google/model"

var (
	RedirectURL    = ""
	DeveloperToken = ""

	AdwordsClientID     = ""
	AdwordsClientSecret = ""

	GaClientID     = ""
	GaClientSecret = ""

	OauthURI = "https://accounts.google.com/o/oauth2/v2/auth?"

	GoogleV9Api = "https://googleads.googleapis.com/v9/customers/%v/googleAds:searchStream"
)
//Init 此处通过资源系统的常用工具token获取 https://res-mgr.mundossp.com/adm/#/admin/account2/313/%E5%85%AC%E7%94%A8%E7%9A%84%E9%85%8D%E7%BD%AE%E8%A1%A8?vID=680&tID
func Init(auth model.Auth) {
	RedirectURL = auth.RedirectURL
	DeveloperToken = auth.DeveloperToken
	AdwordsClientID = auth.AdwordsClientID
	AdwordsClientSecret = auth.AdwordsClientSecret
	GaClientSecret = auth.GaClientSecret
	OauthURI = auth.OauthURI
}
