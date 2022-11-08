package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

const (
	huoziyinyue = "http://127.0.0.1:8100/search"
	quduoduo    = "http://127.0.0.1:8101/search"
	vfine       = "http://127.0.0.1:8102/search"
	audio100    = "http://127.0.0.1:8103/search"
)

var logXml = `
<seelog>
    <outputs formatid="main">
		<rollingfile type="size" filename="#" maxsize="262144000" maxrolls="3" />
        <filter levels="info,debug,critical,error">
            <console />
        </filter>
    </outputs>
    <formats>
        <format id="main" format="%Date/%Time[%File.%Line][%LEV] %Msg%n"/>
    </formats>
</seelog>
`

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	str := muxCurl(string(body))
	data, _ := json.Marshal(str)
	w.Write(data)
	return
}

type Data struct {
	Type    int              `json:"type"`
	Success string           `json:"success"`
	Error   string           `json:"error"`
	Results []map[string]int `json:"results"`
}

func Request(data string, url string) (string, error) {
	log.Println("进来了", data, "url", url)
	jsonStr := []byte(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("http.NewRequest请求地址不正确", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("响应失败", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
func getIntranetIp() string {
	addr, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addr {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
func muxCurl(data string) []map[string]interface{} {
	tasks := []func() (string, error){
		func() (string, error) { return Request(data, quduoduo) },
		func() (string, error) { return Request(data, huoziyinyue) },
		func() (string, error) { return Request(data, vfine) },
		func() (string, error) { return Request(data, audio100) },
	}
	var wg sync.WaitGroup

	var jsdata []map[string]interface{}
	for _, task := range tasks {
		wg.Add(1)
		go func(t func() (string, error)) {
			datafe, err := t()
			if err != nil {
				log.Println("数据为空", datafe)
			}
			{
				log.Println("数据", datafe)
				var m map[string]interface{}
				var data Data
				json.Unmarshal([]byte(datafe), &data)
				js, _ := json.Marshal(data)
				json.Unmarshal(js, &m)
				m["type"] = data.Type
				m["success"] = data.Success
				m["error"] = data.Error
				m["results"] = data.Results
				jsdata = append(jsdata, m)
			}
			wg.Done()
		}(task)
	}
	wg.Wait()
	log.Println("jsdata", jsdata)
	return jsdata
}

func main() {

	log.Println("程序启动")
	localIp := "127.0.0.1"
	log.Println("ip", localIp)
	http.HandleFunc("/", IndexHandler)
	_ = http.ListenAndServe(localIp+":"+"8111", nil)
}
