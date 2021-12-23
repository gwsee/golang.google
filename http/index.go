package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var SaveActionLogAns = true
var SaveActionDir = `/runtime`
var SaveActionLog = true

func init() {
}
func InitHeader(token, managerId string) map[string]string {
	header := make(map[string]string)
	header["Host"] = "googleads.googleapis.com"
	header["User-Agent"] = "curl"
	header["Content-Type"] = "application/json"
	header["Accept"] = "application/json"
	header["Authorization"] = "Bearer " + token // `ya29.a0ARrdaM_UiFSAE2OHzQPFcybLb58oNpDDkMrb6YLl2g8LisRXH5liumd5d7yXzq0oIS-MRCCalYO4ayKc7xDF5cti2cDp3dp0Tbl-QquF1M0t8uvHJwC8t6lvdwjiqrHWnzKzF7cx6-a2QLbtsAEbzVz9B3BBQw`
	header["developer-token"] = DeveloperToken
	if managerId != "" {
		header["login-customer-id"] = strings.Replace(managerId, "-", "", -1)
	}
	return header
}
func Action(method, urls, action string, postData map[string]string, headers map[string]string, ret string) (by []byte, err error) {
	var req *http.Request
	if strings.Contains(action, "json") {
		buf, _ := json.Marshal(postData)
		req, err = http.NewRequest(method, urls, bytes.NewBuffer(buf))
	} else {

		if strings.ToLower(method) == "post" {
			val := url.Values{}
			for k, v := range postData {
				val.Add(k, v)
			}
			req, err = http.NewRequest(method, urls, strings.NewReader(val.Encode()))
		} else {
			if postData != nil {
				val := url.Values{}
				for k, v := range postData {
					val.Add(k, v)
				}
				req, err = http.NewRequest(method, urls+"?"+val.Encode(), strings.NewReader(val.Encode()))
			} else {
				req, err = http.NewRequest(method, urls, nil)
			}

		}
	}
	if err != nil {
		err = fmt.Errorf("http.NewRequest is fail: %v", err.Error())
		return
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do is fail: %v", err.Error())
		return
	}
	by, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("ioutil.ReadAll is fail: %v", err.Error())
		return
	}
	err = handErr(by)
	if SaveActionLog || err != nil {
		_, err1 := saveFile(urls, method, postData, by, ret, err)
		if err1 != nil {
			fmt.Println(err1.Error(), "err1.Error")
		}
	}
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("post google HTTP CODE:" + fmt.Sprint(resp.StatusCode))
		return
	}
	return
}

func handErr(by []byte) (err error) {
	if strings.Contains(string(by), `"errors":`) {
		err = fmt.Errorf("response fail: %v", string(by))
	}
	return
}
