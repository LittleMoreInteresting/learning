package main

import (
	"encoding/json"
	"fmt"
)

type OrderResponse struct {
	Message   string   `json:"message"`
	Status    int64    `json:"status"`
	Timestamp int64    `json:"timestamp"`
	Value     OrderVal `json:"value"`
}
type OrderVal struct {
	Id               int64               `json:"id"`
	OrderCode        string              `json:"orderCode"`
	ProvinceId       int32               `json:"provinceId"`
	CityId           int32               `json:"cityId"`
	CenterId         int32               `json:"centerId"`
	TotalMny         float64             `json:"totalMny"`
	TotalReceivedMny float64             `json:"totalReceivedMny"`
	FirstOrderId     float64             `json:"firstOrderId"`
	OrderCourseDtos  []map[string]string `json:"orderCourseDtos"`
}

func main() {
	str1 := "{\"message\":\"success\",\"value\":{\"orderCode\":\"NBM230215000123\",\"provinceId\":3,\"cityId\":768,\"centerId\":17,\"totalMny\":0.02,\"totalReceivedMny\":0.01,\"status\":0,\"firstOrderId\":\"0\",\"orderCourseDtos\":[{\"courseId\":\"16402943020709866\",\"courseName\":\"\\u5ba1\\u6838\\u8bfe\\u7a0b\\u603b\\u90e8\\u9879\\u76ee\\u90e8\\u4e1a\\u7ee9\\u8bfe\\u7a0b\",\"comeFromOrderEs\":true,\"courseScopes\":[],\"courseDiscountChange\":false,\"newOrderCourse\":false},{\"courseId\":\"16402838330882045\",\"courseName\":\"2022\\u5e74\\u5c71\\u4e1c\\u9752\\u5c9b\\u5e02\\u56fd\\u8003\\u7b14\\u8bd50113\\u7f51\\u8bfe\\u5546\\u54c1\\u73b0\\u6709\\u89c4\\u521901\",\"comeFromOrderEs\":true,\"courseScopes\":[],\"courseDiscountChange\":false,\"newOrderCourse\":false}],\"newStudentName\":\"\\u540c\\u5b665405\",\"orderDiscountChange\":false,\"studentDiscountChange\":false,\"id\":\"16406132168144834\",\"checkDoubleCourseFlag\":true,\"autoCreatOrderFlag\":false,\"isAddExamInfo\":0,\"midAndCenterId\":\"null,17\"},\"timestamp\":1676881198384,\"traceId\":\"\",\"reportError\":0,\"resopnseType\":0,\"status\":1}"
	var data OrderResponse
	err := json.Unmarshal([]byte(str1), &data)
	fmt.Printf("%v:%v \n", data, err)
	/*json_res, _ := json.Marshal(data)
	fmt.Printf("Json:%v\n", string(json_res))*/
}
