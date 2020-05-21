package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	appKey      = `oI8amHfNKhLvO87K`
	accessToken = `dd31ec984cb4a7f9b41a334b053f948`
)

type RespMsg struct {
	TaskID  string `json:"task_id"`
	Result  string `json:"result"`
	Status  int64  `json:"status"`
	Message string `json:"message"`
}

func main() {
	filename := os.Args[1]
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	contentLen := len(contents)

	urlsite := `http://nls-gateway.cn-shanghai.aliyuncs.com/stream/v1/asr?appkey=` + appKey + `&sample_rate=8000`

	client := http.Client{}
	req, err := http.NewRequest("POST", urlsite, bytes.NewReader(contents))
	req.Header.Add(`X-NLS-Token`, accessToken)
	req.Header.Add(`Content-type`, `application/octet-stream`)
	req.Header.Add(`Content-Length`, strconv.Itoa(contentLen))
	req.Header.Add(`Host`, `nls-gateway.cn-shanghai.aliyuncs.com`)

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}

	msg, _ := ioutil.ReadAll(resp.Body)
	text := string(msg)
	fmt.Println(text)

	var res RespMsg
	err = json.Unmarshal(msg, &res)
	if err != nil {
		return
	}
	fmt.Println(res.Result)
}
