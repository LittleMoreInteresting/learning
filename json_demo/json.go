package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Type    string           `json:"type"`
	Success bool             `json:"success"`
	Error   string           `json:"error"`
	Results []map[string]int `json:"results"`
}

func main() {
	str1 := "  {\"success\":true,\"type\":\"quduoduo\",\"error\":null,\"results\":[{\"score\":16,\"id\":941793},{\"score\":16,\"id\":62292},{\"score\":15,\"id\":495671},{\"score\":13,\"id\":1085363},{\"score\":12,\"id\":8995},{\"score\":12,\"id\":654374},{\"score\":12,\"id\":472841},{\"score\":12,\"id\":424171},{\"score\":12,\"id\":111788},{\"score\":12,\"id\":998705},{\"score\":12,\"id\":276020},{\"score\":12,\"id\":109559},{\"score\":12,\"id\":494426},{\"score\":12,\"id\":215740},{\"score\":12,\"id\":633052},{\"score\":12,\"id\":1060317},{\"score\":12,\"id\":548894},{\"score\":12,\"id\":879103},{\"score\":11,\"id\":117860},{\"score\":11,\"id\":959588}]}"
	var data Data
	_ = json.Unmarshal([]byte(str1), &data)
	fmt.Printf("%v\n", data)
	json_res, _ := json.Marshal(data)
	fmt.Printf("Json:%v\n", string(json_res))
}
