package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

<<<<<<< HEAD
type YAMLConfig struct {
	AccessToken string `yaml:"access_token"`
	AppKey      string `yaml:"app_key"`
}

type XMLConfig struct {
	XMLName     xml.Name `xml:"VRServiceEngine"`
	MakeTestLog string   `xml:"MakeTestLog,attr"`
	Debug       string   `xml:"Debug,attr"`
	EngineType  string   `xml:"EngineType,attr"`
	AccessToken string   `xml:"AppId,attr"`
	AppKey      string   `xml:"appKey,attr"`
	Modelss     Models   `xml:"Models"`
}

type Models struct {
	Version string  `xml:"version,attr"`
	Mods    []Model `xml:"Model"`
}

type Model struct {
	Type         string      `xml:"Type,attr"`
	DefaultValue string      `xml:"DefaultValue,attr"`
	DefaultName  string      `xml:"DefaultName,attr"`
	PinYinMatch  string      `xml:"PinYinMatch"`
	Selects      []Selection `xml:"Selection"`
}

type Selection struct {
	Name    string `xml:"Name,attr"`
	Value   string `xml:"Value,attr"`
	KeyWord string `xml:"KeyWord,attr"`
}
=======
const (
	appKey      = `oI8amHfNKhLvO87K`
	accessToken = `dd31ec984cb4a7f9b41a334b053f948`
)
>>>>>>> 5923ef4c3aec32cfbb6486b123ea1229404f7d15

type RespMsg struct {
	TaskID  string `json:"task_id"`
	Result  string `json:"result"`
	Status  int64  `json:"status"`
	Message string `json:"message"`
}

func main() {
	var xmlConf XMLConfig
	xmlConfContents, err := ioutil.ReadFile("conf/VRServiceEngine.xml")
	if err != nil {
		fmt.Println("xml read err")
		return
	}
	err = xml.Unmarshal(xmlConfContents, &xmlConf)
	if err != nil {
		fmt.Println(xmlConf)
		fmt.Println("xml Parse err")
		return
	}
	//fmt.Println(xmlConf)
	//return

	// var conf YAMLConfig
	// confContents, err := ioutil.ReadFile("conf/VRConfig.yaml")
	// if err != nil {
	// 	fmt.Println("yaml read err")
	// 	return
	// }
	// err = yaml.Unmarshal(confContents, &conf)
	// if err != nil {
	// 	fmt.Println("yaml unmarshal err")
	// 	return
	// }
	//fmt.Println(conf.AccessToken)
	//fmt.Println(conf.AppKey)

	filename := os.Args[1]
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	contentLen := len(contents)

	urlsite := `http://nls-gateway.cn-shanghai.aliyuncs.com/stream/v1/asr?appkey=` + xmlConf.AppKey + `&sample_rate=8000`

	client := http.Client{}
	req, err := http.NewRequest("POST", urlsite, bytes.NewReader(contents))
	req.Header.Add(`X-NLS-Token`, xmlConf.AccessToken)
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
		return
	}

	//msg, _ := ioutil.ReadAll(resp.Body)
	//text := string(msg)
	//fmt.Println(text)
	var res RespMsg
	//err = json.Unmarshal(msg, &res)
	//if err != nil {
	//	return
	//}
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}
	fmt.Println(res.Result)
}
