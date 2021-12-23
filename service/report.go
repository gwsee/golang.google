package service

import (
	"encoding/json"
	"fmt"
	"strings"

	"git.zx-tech.net/pengfeng/google/http"
	"git.zx-tech.net/pengfeng/google/model"
)

func GetAdGroupAdList(accountId, managerId, token, start, end string) (list []model.AdGroupAdInfo, err error) {
	var res model.AdGroupAdList
	header := http.InitHeader(token, managerId)
	data := make(map[string]string)
	data["query"] = fmt.Sprintf(sql001, start, end)
	by, err := http.Action("post", fmt.Sprintf(http.GoogleV9Api, strings.Replace(accountId, "-", "", -1)), "json", data, header, "json")
	if err != nil {
		return
	}
	err = json.Unmarshal(by, &res)
	if err != nil {
		return
	}
	for _, v := range res {
		list = append(list, v.Results...)
	}
	return

}
