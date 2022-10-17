package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	huoziyinyue = "http://127.0.0.1:15563/api"
	quduoduo    = "http://127.0.0.1:15563/api"
	vfine       = "http://127.0.0.1:15563/api"
	audio100    = "http://127.0.0.1:15563/api"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	str := muxCurl(string(body))
	data, _ := json.Marshal(str)
	w.Write(data)
	return
}

type Data struct {
	Hash_list []int64 `json:"hash_list"`
	Time_list []int64 `json:"time_list"`
}

func Request(data string, url string) string {
	jsonStr := []byte(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	return string(body)
}
func muxCurl(data string) []map[string]interface{} {
	tasks := []func() string{
		func() string { return Request(data, quduoduo) },
		func() string { return Request(data, huoziyinyue) },
		func() string { return Request(data, vfine) },
		func() string { return Request(data, audio100) },
	}
	var wg sync.WaitGroup

	var jsdata []map[string]interface{}
	for _, task := range tasks {
		task := task
		wg.Add(1)
		go func() {
			datafe := task()
			var m map[string]interface{}
			var data Data
			json.Unmarshal([]byte(datafe), &data)
			js, _ := json.Marshal(data)
			json.Unmarshal(js, &m)
			m["hash_list"] = data.Hash_list
			m["time_list"] = data.Time_list
			jsdata = append(jsdata, m)
			wg.Done()
		}()
	}
	wg.Wait()
	return jsdata
}

func main() {
	http.HandleFunc("/", IndexHandler)
	_ = http.ListenAndServe("127.0.0.1:1234", nil)
}
