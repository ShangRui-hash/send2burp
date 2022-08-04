package burpapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ShangRui-hash/send2burp/debug"
	"github.com/sirupsen/logrus"
)

func Scan(url, configurationName string) (err error) {
	// curl -vgw "\n" -X POST 'http://localhost:1337/v0.1/scan' -d '{"urls":[]}'
	burpSever := "http://localhost:1337/v0.1/"
	burpScanAPI := burpSever + "scan"

	request, err := http.NewRequest("POST", burpScanAPI, nil)
	if err != nil {
		logrus.Error("http.NewRequest failed,err:", err)
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	scanTask := &ScanTask{
		URLs: []string{url},
		ScanConfigurations: []ScanConfiguration{
			{
				Name: configurationName,
				Type: NAMED_CONFIGURATION,
			},
		},
	}

	requestBody, err := json.Marshal(scanTask)
	debug.Println(string(requestBody))
	if err != nil {
		logrus.Error("json.Marshal failed,err:", err)
		return err
	}
	request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
	//发送请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		logrus.Error("http.DefaultClient.Do failed,err:", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		logrus.Info("create scan task success:", url)
	} else {
		logrus.Error("create scan task failed:", url)
	}
	return nil
}
