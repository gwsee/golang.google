package http

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func saveFile(urls, method string, postData map[string]string, by []byte, ret string, errs error) (fileName string, err error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		if errs == nil {
			logFilePath = dir + SaveActionDir + "/google_log/" + now.Format("200601") + "/" + now.Format("02") + "/"
		} else {
			logFilePath = dir + SaveActionDir + "/google_log/" + now.Format("200601") + "/" + now.Format("02") + "_errs/"
		}
	}
	if _, err2 := os.Stat(logFilePath); os.IsNotExist(err2) {
		os.MkdirAll(logFilePath, 0777)
		os.Chmod(logFilePath, 0777)
	}
	{
		// 请求记录
		var saveObj []byte
		saveObj = append(saveObj, []byte(urls)...)
		saveObj = append(saveObj, '\n')
		saveObj = append(saveObj, []byte(method)...)
		saveObj = append(saveObj, '\n')
		buf, _ := json.Marshal(postData)
		saveObj = append(saveObj, buf...)
		saveObj = append(saveObj, '\n')
		if errs != nil {
			saveObj = append(saveObj, []byte(ret)...)
			saveObj = append(saveObj, '\n')
			saveObj = append(saveObj, []byte(errs.Error())...)
			saveObj = append(saveObj, '\n')
		}
		saveObj = append(saveObj, []byte("amazon_request_finished")...)
		saveObj = append(saveObj, '\n')

		logFileName := now.Format("1504") + ".log"
		fileName = logFilePath + logFileName
		fileObj, err1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err1 != nil {
			err = err1
			return
		}
		defer fileObj.Close()
		fileObj.Write(saveObj)
	}
	//保存文件结束！！
	if ret != "" && errs == nil && SaveActionLogAns {
		no := now.Format("1504") + fmt.Sprintf("_%v_%v", time.Now().UnixNano(), rand.Intn(1000))
		{
			logFileName := no + "." + ret
			fileName = logFilePath + logFileName
			fileObj, err1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err1 != nil {
				err = err1
				return
			}
			defer fileObj.Close()
			fileObj.Write(by)
		}
	}
	return
}
